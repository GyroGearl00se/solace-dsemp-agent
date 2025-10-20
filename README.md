# Solace Declarative SEMP Agent (solace-dsemp-agent)

[![Go Version](https://img.shields.io/badge/go-1.24+-blue.svg)](https://golang.org) 
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

**`solace-dsemp-agent`** is a Go application for declaratively managing Solace PubSub+ brokers. It acts as a lightweight agent that reconciles the broker's configuration with a desired target state, embracing GitOps and "Configuration-as-Code" principles for Solace.

The agent listens for a target state definition on a Solace topic. When a new state is published, the agent compares it to the broker's current state and applies the necessary create, update, or delete operations via the Solace SEMPv2 API.

## ⚠️ <span style="color:red">This agent is experimental and subject to change. Use at your own risk!</span>

## 🎯 Key Features

- **Declarative Management**: Define your entire broker configuration (queues, ACLs, client profiles, etc.) in a single JSON file.
- **Event-Driven**: The agent is triggered by publishing a new target state to a Solace topic.
- **Secret Management**: Securely provide secrets like passwords through environment variables or AES-encrypted strings directly in your state file.
- **Selective Control**: Enable or disable management for specific resource types (e.g., only manage queues and ACLs).
- **Status Reporting**: Get feedback on configuration runs via an HTTP webhook or a status message published back to a Solace topic.
- **Dry-Run Mode**: See what changes would be made without actually applying them.

## ⚙️ How It Works

The agent operates in a continuous reconciliation loop:

1.  **Listen**: The agent connects to a Solace broker and subscribes to a configured topic to listen for incoming target state messages.
2.  **Receive State**: A CI/CD pipeline, script, or user publishes a `targetstate.json` message to the topic.
3.  **Fetch Current State**: The agent connects to the broker's SEMP API and fetches the current configuration of all managed resources.
4.  **Compare & Reconcile**: It performs a diff between the desired state (from the message) and the actual state (from the broker).
5.  **Apply Changes**: The agent executes the necessary `CREATE`, `UPDATE`, and `DELETE` operations on the broker to align its configuration with the target state.
6.  **Statusreport**: After the run, it sends a success or failure report to a configured webhook or Solace topic.

### ⚠️ <span style="color:orange">Any managed resource which is not declared in a targetState will be deleted! (Except for System resources having a '#' prefix or equals to 'default'). </span>

If you're not sure about the outcome of your targetState, consider using `SOL_DRYRUN: true` first.



## 🚀 Getting Started

### 1. Configuration

The agent is configured using a `config.yaml` file or environment variables (environment variables always take precedence).

Create a `config.yaml` in the project root. See `config.yaml` for a full example.

**Minimal `config.yaml`:**
```yaml
# Solace SEMP broker configuration (for managing resources)
SOL_SEMP_BROKER_URL: "http://<your-broker-ip>:8080"
SOL_SEMP_USER: "admin"
SOL_SEMP_PASS: "admin"
SOL_SEMP_MSG_VPN: "default"

# Solace messaging configuration (for consuming state)
SOL_STATE_BROKER_URL: "tcp://<your-broker-ip>:55555"
SOL_STATE_MSG_VPN: "default"
SOL_STATE_USERNAME: "my-app-user"
SOL_STATE_PASSWORD: "my-app-password"
SOL_STATE_TOPIC: "config/my-broker/target-state"

# Enable management for specific resources
SOL_MANAGE_QUEUES: true
SOL_MANAGE_ACL_PROFILES: true
```

### 2. 📝 Define Target State

Create a JSON file (`targetstate.json`) that defines the resources you want to manage.

**Example `targetstate.json`:**
```json
{
  "version": "1.0.1",
  "queues": [
    {
      "queueName": "my-app-queue",
      "accessType": "exclusive",
      "permission": "consume",
      "ingressEnabled": true,
      "egressEnabled": true
    }
  ],
  "aclProfiles": [
    {
      "aclProfileName": "my-app-acl",
      "clientConnectDefaultAction": "allow",
      "publishTopicDefaultAction": "disallow",
      "subscribeTopicDefaultAction": "disallow"
    }
  ]
}
```

### 3. 🐳 Run the Agent

A `Dockerimage` is provided to run the agent in a container.

Mount your `config.yaml` file into the container
```bash
docker run --rm -it \
  -v $(pwd)/config.yaml:/app/config.yaml \
  --name solace-dsemp-agent \
  gyrogearl00se/solace-dsemp-agent:latest
```
Alternatively, provide all configuration via environment variables:
```bash
docker run --rm -it \
  -e SOL_SEMP_BROKER_URL="http://<your-broker-ip>:8080" \
  -e SOL_SEMP_USER="admin" \
  -e SOL_SEMP_PASS="admin" \
  -e SOL_DRYRUN="true" \
  # ... other env vars
  --name solace-dsemp-agent \
  gyrogearl00se/solace-dsemp-agent:latest
```


The agent will start and wait for a target state message on the `config/my-broker/target-state` topic.

### 4. 📨 Publish the Target State

Use a tool like `curl` or a Solace client library to publish the content of your `targetstate.json` to the configured topic. The `Makefile` provides a convenience target for this.

```bash
# Make sure the credentials and topic match your config
curl -X POST -d @targetstate.json \
  -u my-app-user:my-app-password \
  http://<your-broker-ip>:9000/TOPIC/config/my-broker/target-state
```

The agent will receive the message and apply the changes.




### 🔐 Secret Management

For sensitive data like passwords in your `targetstate.json`, you can use placeholders that will be resolved by the agent at runtime.

-   **Environment Variable**: `$env{VAR_NAME}`
    
    The agent will replace this with the value of the `VAR_NAME` environment variable.

    ```json
    {
      "clientUsernames": [
        {
          "clientUsername": "my-secure-user",
          "password": "$env{USER_PASSWORD}"
        }
      ]
    }
    ```

-   **AES Encryption**: `$aes{BASE64_ENCRYPTED_VALUE}`

    The agent will decrypt the value using the AES key provided in the `SOL_AES_KEY` configuration variable.
    
    - Encryption Key Size: 16, 24, or 32 bytes (AES-128, AES-192, or AES-256)
    - Encryption Mode: CBC
    - Output format: Base64

    ```json
    {
      "clientUsernames": [
        {
          "clientUsername": "my-secure-user",
          "password": "$aes{mHzBHwp+08JjXnLLfmHyZ5Acw0ecTshvcLoW7AABgdg=}"
        }
      ]
    }
    ```

### 💼 Supported Solace Objects

The agent supports managing a wide range of Solace resources:
- `msgVpns`
- `queues` (including `queueSubscriptions`)
- `aclProfiles` (including `publishExceptions` and `subscribeExceptions`)
- `clientUsernames`
- `clientProfiles`
- `bridges`
- `bridgeRemoteMsgVpns`
- `dmrBridges`
- `jndiConnectionFactories`
- `jndiQueues`
- `jndiTopics`
- `proxies`
- `queueTemplates`
- `topicEndpoints`
- `topicEndpointTemplates`


### 📊 Statusreport
The statusreport (Webhook or via Messaging) is a structured json containing following informations:
- all additional fields provided via SOL_STATUS_EXTRA_FIELDS
- success status
- errors summary
- "version" field provided in the targetstate.json
- timestamp of the run

## 🛠️ Development

This project uses a [Devcontainer](.devcontainer/) to ensure a consistent development environment.

### 🧡 Contributing
Contributions are welcome!

- Fork this repository.
- Make changes within your fork.
- Create a pull request.

## ⚖️ License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.