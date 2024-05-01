# Env Loader

**Env-Loader is a tool that loads and parses environment variables into a struct.**

## Motivations

If you're like me, you wouldn't like to write multiple `os.LookupEnv()` multiple times. Even if you have no problem with writing multiple `os.LookupEnv()`, you still have to parse strings into respective types.

## Solution

Intorducing env-loader. Now you can just declare a struct type with the name of the environment variable you want and pass it to a function. Then magic happens. It will loads your environment variables, parses those into respective types and set them to the struct you provided.

> You can load all of your environment variables in one struct or you can separate them into multiple structs. (This is my favorite feature.)
