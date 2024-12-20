# docker-plugin-cinder

This Docker volume plugin for utilizing OpenStack Cinder for persistent storage volumes.

The plugin attaches block storage volumes to the compute instance running the plugin. If the volume is already attached to another compute instance it will be detached first.

## Requirements

* Block Storage API v3
* Compute API v2
* KVM w/ virtio or SCSI disks

## Install

Download binary or package from the latest release. The packages include a systemd service and socket file, as well as sample config file.

### Configuration

Provide configuration at `/etc/docker/cinder.json` for the plugin:

```json
{
    "endpoint": "http://keystone.example.org/v3",
    "username": "username",
    "password": "password",
    "domainID": "",
    "domainName": "default",
    "tenantID": "",
    "tenantName": "",
    "applicationCredentialId": "",
    "applicationCredentialName": "",
    "applicationCredentialSecret": "",
    "region": "",
    "mountDir": ""
}
```

Enable and start the `docker-plugin-cinder.socket` unit and start/restart docker:

```console
systemctl enable docker-plugin-cinder.socket
systemctl start docker-plugin-cinder.socket

systemctl restart docker.service
```

Or manually run the plugin before docker:

```console
$ /usr/bin/docker-plugin-cinder -config /etc/docker/cinder.json
INFO Connecting...                                 endpoint="http://api.os.xopic.de:5000/v3"
INFO Machine ID detected                           id=e0f89b1b-ceeb-4ec5-b8f1-1b9c274f8e7b
INFO Connected.                                    endpoint="http://api.os.xopic.de:5000/v3"
```

By default, a `cinder.json` from the current working directory will be used.

## Usage

The default volume size is 10 GB but can be overridden:

```console
docker volume create -d cinder -o size=20 volname
```

## Notes

### Machine ID

This plugin expects `/etc/machine-id` to be the OpenStack compute instance UUID which seems to be the case when booting cloud images with KVM. Otherwise, configure `machineID` in the configuration file.

### Attaching volumes

Requested volumes that are already attached will be forcefully detached and moved to the requesting machine.

## License

MIT License
