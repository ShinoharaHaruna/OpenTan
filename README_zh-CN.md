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
> OpenTan 对因使用 OpenTan 造成的任何后果概不负责。

## 使用方法

在设置好 Go 环境并克隆存储库（或下载二进制文件）后，您就可以使用 OpenTan 了。

### 配置

将 `config.template.yml` 文件复制到 `config.yml`。您可以设置运行的主机和端口，并将 `Mode` 更改为 `debug` 以获取更多信息。

`Prefix` 字段用于设置 API 端点的前缀。默认值为 `/v1`，因此 API 端点默认为 `http://localhost:52711/v1`。

`ID` 和 `Password` 用于保持您的令牌新鲜。它们与您登录 MyTan 时使用的完全相同。通常来说，`ID` 是您注册的手机号。

从逻辑上讲，`API_Key` 不是必需的，因为 OpenTan 会自动为您生成一个。但是，您必须将其保存在配置文件中，以确保 OpenTan 的功能正常。因此，您可以填入任何内容以开始使用 OpenTan。

### Cline 集成

我们假设您已掌握 Go 和 Cline 的基本知识，并且已经运行了 OpenTan 并在 VS Code 中安装了 Cline。

1. 打开 Cline 设置。
2. 选择 `OpenAI Compatible` 作为 API 提供程序。
3. 设置 API 端点，例如 `http://localhost:52711/v1`。
4. 在 `API_Key` 中键入任何内容，因为 OpenTan 具有自己的 API 逻辑。
5. 在 `Model ID` 中键入任何内容，因为 OpenTan 具有自己的模型设置。您可以通过配置文件指定使用的模型，而默认配置使得 OpenTan 使用 `gemini-2.5-flash` 作为默认模型。
6. 现在您可以将 Cline 与 OpenTan 一起使用了！如果您有其他设置，请在 Cline 中进行所需的操作。

## 模型

和 OpenTan 工作最好的模型是 `gemini-2.5-flash`，它在处理长提示时表现良好；如果使用 `mai-seed-pro` 等模型，可能会面临一些问题。

以下是一些推荐使用的模型：

- `gpt-4o`：已经不错。无消耗。
- `gpt-4o-mini`：可以满足大多数需求，无消耗。
- `gpt-4.1`：大概会更好一些。正常一倍率消耗。
- `gemini-2.5-flash`：最佳实践，兼顾长上下文和性能。无消耗。
- `mai-seed-pro`：基于 `Deepseek-V3-0324`，总体体验优于 `gpt-4o`。无消耗。
- `mai-seed-think`：基于 `Deepseek-R1`，感觉太重了，实际体验上来说不适合稍复杂的场景，会导致思考时间过长。无消耗。

## 已知问题

> 更准确地说，是特性（

OpenTan 会在每次请求后清除历史记录，使得 Cline 产生的对话不会影响用户在 MyTan 上正常的历史记录。您会在「最近删除」列表中看到它们。

---

当与 Cline 一起使用时，它很容易生成极长的 prompt，MytTan 可能无法正确响应。在这种情况下，我个人的最佳实践是放弃原来的会话，并对新会话使用相同的任务提示。

## 贡献

我们欢迎任何形式的贡献！如果您想为 OpenTan 做出贡献，请随时提出 issue 或 pull request。

## 许可

OpenTan 在 MIT 许可下获得许可。有关详细信息，请参见 [LICENSE](LICENSE) 文件。

OpenTan 与 MyTan 或 OpenAI 没有任何关系，并且对因使用 OpenTan 造成的任何后果概不负责。
