package constant

// Error represents the different error messages used in the application.
type Error string

// Constants for various error messages that can occur in the application.
const (
	
	FailedToParseData            	Error = "Oops! We encountered an issue while processing your data. Please try again!"
	FailedToInsertData           	Error = "Sorry, we couldn’t add the new data. Please try again!"
	FailedToUpdateData           	Error = "We couldn’t update the data. Please give it another shot!"
	FailedToDeleteData           	Error = "We couldn’t delete the data. Please try again!"
	FailedToRetrieveData         	Error = "There was an issue retrieving your data. Please try again!"
	FailedToHashPassword         	Error = "We ran into an error while securing your password. Please try again!"
	FailedToGenerateTokenAccess  	Error = "We couldn’t generate an access token. Please try again!"
	FailedToLoadUserData         	Error = "We couldn’t load your account information. Please try again!"
	FailedToDecodeData           	Error = "We couldn’t process the data. Please try again!"
	FailedToExchangeCodeWithToken 	Error = "Authentication couldn’t be completed. Please try again!"
	FailedToGetCodeFromRedirectUrl 	Error = "We couldn’t retrieve the authorization code. Please try again!"
	FailedToLoadMessage          	Error = "Failed to load your message. Please try again!"
	FailedToSendMessage          	Error = "Failed to send your message. Check your connection and try again!"
	
	DuplicateDataError           	Error = "This username or email is already taken. Please choose another one!"
	UnregisteredAccountError     	Error = "It seems like this email isn’t registered. Please sign up to continue!"
	InvalidAccountError          	Error = "The email or password entered is incorrect. Please try again!"
	InvalidIdError               	Error = "The provided ID is invalid. Please double-check and try again!"
	ValidationError              	Error = "Your input doesn’t meet the requirements. Please check and try again!"
	InvalidTokenError            	Error = "The token provided is invalid. Please try again!"
	PermissionDeniedError        	Error = "Access denied! You don’t have permission to access this data."
	DataUnavailableError         	Error = "This data is currently unavailable. Please try again later!"
	GreetingEmailError           	Error = "Unable to send greeting email due to an unexpected error."

)
