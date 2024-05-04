# Env Loader

**Env-Loader is a tool that loads and parses environment variables into a struct.**

## Installation

```bash
go get github.com/Shin-Thant/env-loader
```

## Motivations

If you're like me, you wouldn't like to write multiple `os.LookupEnv()` multiple times. Even if you have no problem with writing multiple `os.LookupEnv()`, you still have to parse strings into respective types.

## Solution

Intorducing env-loader! Now you can just declare a struct type with the name of the environment variable you want and pass it to a function. Then magic happens. It will loads your environment variables, parses those into respective types and set them to the struct you provided.

> You can load all of your environment variables in one struct or you can separate them into multiple structs. (This is my favorite feature.)

## Usage Guide

1. When your `.env` file is in your project root directory, you can just call like this.

    ```golang
    type appEnv struct {
        PORT int
    }
    envConfig := appEnv
    envloader.LoadEnv(&envConfig, nil)
    ```

2. If your `.env` file is in different directory or you have different name for example, `.env.development`, you must specify the **_relative path_** from the root directory.

    ```golang
    type appEnv struct {
        PORT int
    }
    envConfig := appEnv
    envloader.LoadEnv(&envConfig, &envloader.LoadEnvOptions{
        EnvPath: "/.env.development",
    })
    ```

## Docker

```bash
docker run --name envloader_test -it --rm shinthant101/env-loader:main
```
