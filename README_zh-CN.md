# OpenTan

<div align="right">
<img src="./icon.jpg" width="100" align="right" />
</div>
<img src="https://img.shields.io/github/license/ShinoharaHaruna/OpenTan" />
<img src="https://img.shields.io/github/go-mod/go-version/ShinoharaHaruna/OpenTan" />

[English README](README.md)

## 简介

OpenTan 是一个 MyTan API 包装器，旨在与 OpenAI API 兼容。更具体地说，它模拟了 `/chat/completions` 端点。通过运行 OpenTan，您可以将其与 Cline 结合使用，以极低的成本获得 AI IDE 体验。

> **警告：** 使用 OpenTan **可能** 违反 MyTan 的服务条款，或损害 MyTan 服务器。正常使用 OpenTan 是安全的，但请怀着善意使用，不要滥用。
>
> **强烈不建议** 将 Anthropic Claude 的高端模型与 OpenTan 一起使用，因为这可能会对 MyTan 服务器造成巨大损害。
>
> 强烈建议对 OpenTan 的 RateLimiter 保持保守配置。
>
> OpenTan 对因使用 OpenTan 造成的任何后果概不负责。

## 使用方法

在设置好 Go 环境并克隆存储库（或下载二进制文件）后，您就可以使用 OpenTan 了。

### 配置

将 `config.template.yml` 文件复制到 `config.yml`。您可以设置运行的主机和端口，并将 `Mode` 更改为 `debug` 以获取更多信息。

`Prefix` 字段用于设置 API 端点的前缀。默认值为 `/v1`，因此 API 端点为 `http://localhost:52711/v1`。

`ID` 和 `Password` 用于保持您的令牌新鲜。它们与您登录 MyTan 时使用的完全相同。

从逻辑上讲，`API_Key` 不是必需的，因为 OpenTan 会自动为您生成一个。但是，您必须将其保存在配置文件中，以确保 OpenTan 的功能正常。

`RateLimit` 部分中的所有字段均以秒为单位。`Rate` 定义每秒允许的请求数，`Burst` 定义令牌桶的大小。`MaxWait` 是请求的最大等待时间。我们 **不建议** 将这些值设置为较高的数字，因为它可能会损害 MyTan。

### Cline 集成

我们假设您已掌握 Go 和 Cline 的基本知识，并且已经运行了 OpenTan 并在 VS Code 中安装了 Cline。

1. 打开 Cline 设置。
2. 选择 `OpenAI Compatible` 作为 API 提供程序。
3. 设置 API 端点，例如 `http://localhost:52711/v1`。
4. 在 `API_Key` 中键入任何内容，因为 OpenTan 具有自己的 API 逻辑。
5. 在 `Model ID` 中键入任何内容，因为 OpenTan 具有自己的模型设置。OpenTan 使用 `gemini-2.0-flash` 作为默认模型，目前您只能在源代码中修改它。但是，我们建议您在 Cline 的 `Model ID` 字段中键入它，以便更好地识别。
6. 现在您可以将 Cline 与 OpenTan 一起使用了！如果您有其他设置，请在 Cline 中进行所需的操作。

## 已知问题

> 更准确地说，是特性（

OpenTan 模拟了 MyTan 的正常使用，因此会在您的历史记录列表中创建大量的对话。因此，如果您想保持 MyTan 帐户历史记录的整洁，请非常保守地使用 OpenTan。

## 贡献

我们欢迎任何形式的贡献！如果您想为 OpenTan 做出贡献，请随时提出 issue 或 pull request。

## 许可

OpenTan 在 MIT 许可下获得许可。有关详细信息，请参见 [LICENSE](LICENSE) 文件。

OpenTan 与 MyTan 或 OpenAI 没有任何关系，并且对因使用 OpenTan 造成的任何后果概不负责。
