package usermodel

type Users struct {
	ID	int
	CODE	string	
	CREATED_AT string
	CREATED_BY	int
	IS_DELETED	int
	UPDATED_AT	string	
	UPDATED_BY	int	
	EMAIL	string	
	IS_ACTIVE	int	
	IS_DELEGATION	int	
	IS_LOCKED		int	
	LOGIN_ATTEMPT	int	
	NAME	string	
	PASSWORD	string	
	PROPOSE_DATE	string	
	PROPOSE_TOKEN	string	
	PROPOSE_TYPE	string	
	PROPOSE_VALUE	string	
	ORGANIZATION_ID	int	
	IMG_FILE_ID		int	
	LANG_ID		int	
	LOG_DATE	string	
	LOGIN_ATTEMPT_DATE	string	
	NIP	string	
	PADI_ORGANIZATION_ID	int	
	USER_AUCTION_CODE		string	
	ROLE_CODE	string	
	ROLE_ID		int	
	ROLE_NAME	string	
}