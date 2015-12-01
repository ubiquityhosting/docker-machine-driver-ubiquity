# Ubiquity Hosting driver for Docker Machine

![](/docs/img/ubiquityhosting.png)

Install this driver in your PATH and you can create docker hosts on [Ubiquity Hosting](https://www.ubiquityhosting.com).

## Installation

Official release versions of the driver include a binary for Linux,
MacOS and Windows. You can find them on the [GitHub releases
page](https://github.com/ubiquity/docker-machine-driver-ubiquity/releases).

Pick the binary you require, download it into a directory on your
PATH as a file called `docker-machine-driver-ubiquity` and make it
executable.

Linux:

```
curl -sSL -o /usr/local/bin/docker-machine-driver-ubiquity \
https://github.com/ubiquityhosting/docker-machine-driver-ubiquity/releases/download/v0.0.1/docker-machine-driver-ubiquity_linux-amd64 && \
chmod 755 /usr/local/bin/docker-machine-driver-ubiquity

```

Mac OSX

```
sudo curl -sSL -o /usr/local/bin/docker-machine-driver-ubiquity https://github.com/ubiquityhosting/docker-machine-driver-ubiquity/releases/download/v0.0.1/docker-machine-driver-ubiquity_darwin-amd64 &&
sudo chmod 755 /usr/local/bin/docker-machine-driver-ubiquity
```

## Obtaining credentials

Login to Ubiquity Motion and navigate to [API Tools](https://motion.ubiquityhosting.com/api).
Take note of your 'Reseller ID', 'Remote ID', and 'Access Key' - these will be needed later.

## Using the driver

To use the driver first make sure you are running at least [version
0.5.1 of `docker-machine`](https://github.com/docker/machine/releases).

```
$ docker-machine -v
docker-machine version 0.5.1 (7e8e38e)
```

Check that `docker-machine` can see the Ubiquity Hosting driver by asking for the driver help.

```
$ docker-machine create -d ubiquity | more
Usage: docker-machine create [OPTIONS] [arg...]

## Create a machine.

Specify a driver with --driver to include the create flags for that driver in the help text.

Options:

| Option                  |  Environment Variable | Default | Description                                     | Required? |
|-------------------------|:---------------------:|---------|-------------------------------------------------|:---------:|
| --ubiquity-api-token    |   UBIQUITY_API_TOKEN  |         | Ubiquity Access Key for authentication          |     Y     |
| --ubiquity-api-username | UBIQUITY_API_USERNAME |         | Ubiquity Remote ID for authentication           |     Y     |
| --ubiquity-client-id    |   UBIQUITY_CLIENT_ID  |         | Ubiquity Reseller ID for account authentication |     Y     |
| --ubiquity-flavor-id    |   UBIQUITY_FLAVOR_ID  |    1    | Ubiquity VM size details for VM creation        |           |
| --ubiquity-image-id     |   UBIQUITY_IMAGE_ID   |    18   | Ubiquity VM image for VM creation               |           |
| --ubiquity-zone-id      |    UBIQUITY_ZONE_ID   |    7    | Ubiquity zone location for VM creation          |           |

...
```

To create a machine you'll need the API credential details you obtained earlier. 

Then creating a Docker host is as simple as

```
$ docker-machine create -d ubiquity --ubiquity-client-id 1234 --ubiquity-api-username ubic-1234 --ubiquity-api-token 273182f1237b361cd9de8b3ea651905d  example
Running pre-create checks...
Creating machine...
Waiting for machine to be running, this may take a few minutes...
Machine is running, waiting for SSH to be available...
Detecting operating system of created instance...
Provisioning created instance...
Copying certs to the local machine directory...
Copying certs to the remote machine...
Setting Docker configuration on the remote daemon...
To see how to connect Docker to this machine, run: docker-machine env example
```

## Changing the settings

The driver has several options that you can use to get precisely the
Docker host you want. You can see them all in the help list by running
`docker-machine create -d ubiquity | more`

## Help

If you need help using this driver, drop an email to support at ubiquityhosting dot com.

## License

This code is released under the MIT License.

Copyright (c) 2015 Ubiquity Hosting, Nobis Technology Group, LLC