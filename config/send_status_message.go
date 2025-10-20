package config

import (
	"solace.dev/go/messaging"
	"solace.dev/go/messaging/pkg/solace/config"
	"solace.dev/go/messaging/pkg/solace/resource"
)

func SendStatusMessage(success bool, payload []byte, brokerURL, topic, brokerUser, brokerPass, msgVpn string) error {

	brokerConfig := config.ServicePropertyMap{
		config.TransportLayerPropertyHost:                brokerURL,
		config.ServicePropertyVPNName:                    msgVpn,
		config.AuthenticationPropertySchemeBasicPassword: brokerPass,
		config.AuthenticationPropertySchemeBasicUserName: brokerUser,
	}

	messagingService, err := messaging.NewMessagingServiceBuilder().FromConfigurationProvider(brokerConfig).Build()

	if err != nil {
		panic(err)
	}

	if err := messagingService.Connect(); err != nil {
		panic(err)
	}

	directPublisher, builderErr := messagingService.CreateDirectMessagePublisherBuilder().Build()
	if builderErr != nil {
		panic(builderErr)
	}

	startErr := directPublisher.Start()
	if startErr != nil {
		panic(startErr)
	}

	messageBody := payload
	messageBuilder := messagingService.MessageBuilder().
		WithProperty("application", "solace-dsemp-agent").
		WithProperty("language", "go")

	message, err := messageBuilder.BuildWithByteArrayPayload(messageBody)
	if err != nil {
		panic(err)
	}
	resourceTopic := resource.TopicOf(topic)

	publishErr := directPublisher.Publish(message, resourceTopic)
	if publishErr != nil {
		panic(publishErr)
	}

	// fmt.Printf("Published message: %s\n", message)

	messagingService.Disconnect()

	return nil

}
