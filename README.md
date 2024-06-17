# go-viper-cobra-skeleton

This is a Golang CLI skeleton project that clones itself to provide initial boilerplate that wires up [Cobra](https://github.com/spf13/cobra), [Viper](https://github.com/spf13/viper) and a logging system, an initial `version` command, along with `Makefile`, `README.MD`, and `LICENSE` files to bootstrap a new project. After the project is generated, a local git repository is initialized along with a first commit.

The generated project will accept its configuration from CLI arguments or any source supported by Viper.

# Build

This project includes a `Makefile` with targets to build the following architectures:

- x86
- x86_64
- ARM
- ARM_64


# Usage

Use the `clone` target to create a new project:

    # Create new project
    CLONE_HOST=bitbucket.org CLONE_OWNER=myorg CLONE_APP_NAME=my-new-app make clone

    # Try it out ...
    cd ../go-my-new-app
    make
    ./my-new-app version --debug

    # That's it, get to work!

The following command results in a new directory being created one level up `../my-new-app` that contains the files included in this skeleton project, with import statements using `bitbucket.org/myorg/my-new-app`


# Generated Skeleton Project

The skeleton project provides initial boilerplate that wires up [Cobra](https://github.com/spf13/cobra), [Viper](https://github.com/spf13/viper) and a logging system, an initial `version` command, and `Makefile`, `README.MD`, and `LICENSE` files.

## Configuration

Viper is initialized to read `config.yaml` from the following locations, along with the directory specified by the `--config-dir` CLI option.

- ./config.yaml
- /etc/{CLONE_APP_NAME}/config.yaml
- $HOME/.{CLONE_APP_NAME}/config.yaml

The following code is used to initialize Cobra:

```
viper.SetConfigName("config")
viper.SetConfigType("yaml")
viper.AddConfigPath(app.ConfigDir)
viper.AddConfigPath(fmt.Sprintf("/etc/%s/", Name))
viper.AddConfigPath(fmt.Sprintf("$HOME/.%s/", Name))
viper.AddConfigPath(".")
```

## Commands

#### version

The included `Makefile` embeds git and build information into the built binary. The `version` command shows the application name, version, and build information.

## Debugging

The skeleton project includes support for a `--debug` flag that outputs the Viper configuration to the console and puts the logging system into "debug mode" so `logger.Debug` statements are sent to the log file and stdout.

## Versioning

The generated skeleton project includes a semantic versioning file in alignment with the [Golang Module version number](https://go.dev/doc/modules/version-numbers) intended for use by your favorite CI/CD system.

## Logging

The logging system sends logs to stdout and the directory specified by the `--log-dir` CLI option, or `log-dir` configuration option set.
