package auth

type RefreshRequestDTO struct {
	RefreshToken string `json:"refreshToken"`
}

type RefreshResponseDTO struct {
	AccessToken string `json:"accessToken"`
}
