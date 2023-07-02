package domain

type MetadataResponse struct {
	Total        int  `json:"total"`
	CurrentPage  int  `json:"current_page"`
	Page         int  `json:"page"`
	HasPrevious  bool `json:"has_previous"`
	HasNext      bool `json:"has_next"`
	PreviousPage int  `json:"previous_page"`
	NextPage     int  `json:"next_page"`
}

type Response struct {
	Success   bool        `json:"success"`
	Message   string      `json:"message"`
	ErrorCode int16       `json:"error_code"`
	Data      interface{} `json:"data"`
}
