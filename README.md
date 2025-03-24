# OpenTan

<div align="right">
<img src="./icon.jpg" width="100" align="right" />
</div>
<img src="https://img.shields.io/github/license/ShinoharaHaruna/OpenTan" />
<img src="https://img.shields.io/github/go-mod/go-version/ShinoharaHaruna/OpenTan" />

[Chinese Documentation](README_zh-CN.md)

## Description

OpenTan is a MyTan API Wrapper designed to be compatible with the OpenAI API. More specifically, it emulates the `/chat/completions` endpoint. By running OpenTan, you can use it with Cline to get an AI IDE experience at a very low cost.

> **Warning:** Using OpenTan **may** violate the terms of service of MyTan or cause harm to MyTan servers. Using OpenTan properly is normally safe, but please use it with a kind heart and do not abuse it.
>
> **Strongly discouraged:** Using high-end models from Anthropic Claude with OpenTan is **strongly discouraged**, as it could potentially cause significant harm to MyTan servers.
>
> It is highly recommended to maintain a conservative configuration for OpenTan's RateLimiter.
>
> OpenTan is not responsible for any consequences caused by using OpenTan.

## Usage

After setting up your Go environment and cloning the repository (or downloading the binary), you are ready to use OpenTan.

### Configuration

Copy the `config.template.yml` file to `config.yml`. You can configure the running host and port, and change `Mode` to `debug` for more detailed information.

The `Prefix` field is used to set the prefix for the API endpoint. The default value is `/v1`, so the API endpoint will be `http://localhost:52711/v1`.

`ID` and `Password` are used to keep your token fresh. These are the same credentials you use to log in to MyTan.

Logically speaking, the `API_Key` is not strictly necessary, as OpenTan will automatically generate one for you. However, you *must* keep it in your configuration file to ensure OpenTan functions correctly.

All fields in the `RateLimit` section are in seconds. `Rate` defines the number of requests allowed per second, and `Burst` defines the size of the token bucket. `MaxWait` is the maximum waiting time for a request. We **strongly advise against** setting these values too high, as it could harm MyTan.

### Cline Integration

We assume you have a basic understanding of Go and Cline, and that you have already run OpenTan and installed Cline in VS Code.

1.  Open the Cline settings.
2.  Choose `OpenAI Compatible` as the API Provider.
3.  Set the API Endpoint, e.g., `http://localhost:52711/v1`.
4.  Type anything in `API_Key`, as OpenTan has its own API logic.
5.  Type anything in `Model ID`, as OpenTan has its own model settings. OpenTan uses `gemini-2.0-flash` as the default model, and currently, you can only modify this in the source code. However, we recommend typing it in the `Model ID` field of Cline for better recognition.
6.  Now you can use Cline with OpenTan! Configure any further settings you desire within Cline.

## Known Issues

> Or rather, features（

OpenTan emulates the normal usage of MyTan, so it creates tons of conversations in your history list. So if you want to keep history of your MyTan account clean, use OpenTan with great conservatism.

When used with Cline, which can easily generate extremely long prompts, MytTan may not respond properly.

## Contributing

We welcome contributions of any kind! If you would like to contribute to OpenTan, please feel free to open an issue or a pull request.

## License

OpenTan is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

OpenTan has NO affiliation with MyTan or OpenAI, and holds no responsibility for any consequences resulting from the use of OpenTan.
