// internal/global/model/modelInfo.go

package model

type ModelInfoResponse struct {
    Success bool          `json:"success"`
    Errors  []interface{} `json:"errors"` // Assuming errors are interface{} type
    Data    ModelData     `json:"data"`
}

type ModelData struct {
    Models []ModelInfo `json:"models"`
}

type ModelInfo struct {
    ModelName           string        `json:"model_name"`
    Description         string        `json:"description,omitempty"`
    MaxTokens           int           `json:"max_tokens"`
    Status              string        `json:"status"`
    Order               int           `json:"order"`
    UserLimit           int           `json:"user_limit"`
    UserLimitDuration   int           `json:"user_limit_duration"`
    UserMaxTokens       int           `json:"user_max_tokens"`
    UserMinTokens       int           `json:"user_min_tokens"`
    UserDefaultTokens   int           `json:"user_default_tokens"`
    PlusOnly            bool          `json:"plus_only,omitempty"`
    PromptTokenRate     float64       `json:"prompt_token_rate,omitempty"`
    CompletionTokenRate float64       `json:"completion_token_rate,omitempty"`
    PluginConfig        *PluginConfig `json:"plugin_config,omitempty"`      // Using pointer to handle null values
    IsVoiceAvailable    *bool         `json:"is_voice_available,omitempty"` // Using pointer to handle null values
    ImageConfig         *ImageConfig  `json:"image_config,omitempty"`       // Using pointer to handle null values
    Icons               Icons         `json:"icons"`
    UserLimitTTL        int           `json:"user_limit_ttl"`
}

type PluginConfig struct {
    MaxAllowedPlugins    int      `json:"max_allowed_plugins"`
    PluginConfigComment  string   `json:"plugin_config_comment"`
    OptionalPluginIDs    []string `json:"optional_plugin_ids"`
    DefaultSystemPrompt  string   `json:"default_system_prompt"`
    DefaultContextLength int      `json:"default_context_length"`
}

type ImageConfig struct {
    IsImageRecognitionAvailable bool `json:"is_image_recognition_available"`
    IsTextRequired              bool `json:"is_text_required"`
    MaxUploadImage              int  `json:"max_upload_image"`
}

type Icons struct {
    Colorful         string `json:"colorful,omitempty"`
    Colorless        string `json:"colorless,omitempty"`
    IconColorfulApp  string `json:"icon_colorful_app"`
    IconColorlessApp string `json:"icon_colorless_app"`
    IconColorfulWeb  string `json:"icon_colorful_web"`
    IconColorlessWeb string `json:"icon_colorless_web"`
    AvatarApp        string `json:"avatar_app"`
    AvatarWeb        string `json:"avatar_web"`
}
