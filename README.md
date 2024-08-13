# Chat Message Broker
  The Chat Message Broker is a distributed server application that manages communication between clients. It includes both server (broker) and client applications.
# Broker
  The broker application accepts the following flags:

  --config: Specifies the path to the configuration file. Default value is ./broker/config.yaml. Modify this path if you use a different location for the config file.
  --connection: Specifies the socket connection type. Supported values are tcp-socket and websocket. The default value is all, which enables both connection types.
# Client
  The client application accepts the following flags:

  --host: Specifies the broker host address. Default value is 0.0.0.0.
  --port: Specifies the broker host port number. Default value is 8888.
  --connection-type: Specifies the broker connection type. Default value is tcp.
Modify these values only if you need to override the defaults.

# Usage
  The broker accepts three types of messages:

  Direct Message: Begins with @. Example: @{$USER_NAME} {$MESSAGE_TEXT}
  Channel Message: Begins with #. Example: #{$CHANNEL_NAME} {$MESSAGE_TEXT}
  Command Message: Begins with /. Example: /{$SUPPORTED_COMMAND}
  At startup, the broker creates a default channel named random.

  When a client connects, the broker registers the new user, who is automatically subscribed to the random channel.

# Supported Commands
  The broker supports various commands to interact with the workspace:

  /create: Creates a new channel. The user who creates the channel is subscribed to it by default. Example: /create {$CHANNEL_NAME}
  /join: Joins a specific channel. Example: /join {$CHANNEL_NAME}
  /leave: Leaves a specific channel. Example: /leave {$CHANNEL_NAME}
  /show: Displays information about the workspace or channels. Example: /show {$OPTION}

# The show command supports the following options:

users: Displays a list of all workspace users.
channels: Displays a list of all workspace channels.
all: Displays a list of all workspace users and channels.
$CHANNEL_NAME: Displays a list of all users subscribed to a specific channel. Example: /show random
