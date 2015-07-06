# DaemonGarden

> A daemon to run others - controllable through a REST interface

## Usage

```
$ daemon-garden --help

Usage of bin/daemon-garden:
  -address="127.0.0.1:8081": address to bind host:port
  -logDir="/var/log/daemonGarden": directory to put my daemon logs to
```

## Testing

```
$ git clone https://github.com/foomo/DaemonGarden.git
$ cd content-server
$ make test
```

## Contributing
In lieu of a formal styleguide, take care to maintain the existing coding style. Add unit tests and examples for any new or changed functionality.

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request

## License
Copyright (c) foomo under the LGPL 3.0 license.
