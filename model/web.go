package model

type (
	UserCreateRequest struct {
		Username             string `json:"username"`
		Email                string `json:"email"`
		Handphone            string `json:"handphone"`
		Password             string `json:"password"`
		PasswordConfirmation string `json:"password_confirmation"`
	}

	UserUpdateRequest struct {
		ID        string
		Username  string `json:"username"`
		Email     string `json:"email"`
		Handphone string `json:"handphone"`
	}

	UserResponse struct {
		ID        string `json:"id"`
		Username  string `json:"username"`
		Email     string `json:"email"`
		Handphone string `json:"handphone"`
	}

	KelontongCreateRequest struct {
		Name     string `json:"name"`
		Location string `json:"location"`
	}

	KelontongUpdateRequest struct {
		ID       string
		Name     string `json:"name"`
		Location string `json:"location"`
	}

	KelontongResponse struct {
		ID       string `json:"id"`
		Name     string `json:"name"`
		Location string `json:"location"`
	}

	FashionCreateRequest struct {
		Name string `json:"name"`
	}

	FashionUpdateRequest struct {
		ID   string
		Name string `json:"name"`
	}

	FashionResponse struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}

	WebResponse struct {
		Code   int         `json:"code"`
		Status string      `json:"status"`
		Data   interface{} `json:"data"`
		Error  interface{} `json:"errors"`
	}
)
