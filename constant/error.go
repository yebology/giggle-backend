package constant

type Error string

const (

	FailedToParseData Error = "There was an error while processing the data!"
	FailedToInsertData Error = "Failed to add new data. Please try again!"
	FailedToUpdateData Error = "Failed to update the data. Please try again!"
	FailedToDeleteData Error = "Failed to delete the daa. Please try again!"
	FailedToRetrieveData Error = "There was an issue retrieving the data. Please try again!"
	FailedToHashPassword Error = "There was an error while securing your password!"
	FailedToGenerateTokenAccess Error = "Unable to generate access token. Please try again!"
	FailedToLoadUserData Error = "Failed to load your account information. Please try again!"
	FailedToDecodeData Error = "We couldn't process the data. Please try again!"

	DuplicateDataError Error = "This username or email is already taken. Please choose another one!"
	UnregisteredAccountError Error = "This email address is not registered. Please check your email or sign up!"
	InvalidAccountError Error = "The email or password you entered is incorrect. Please try again!"
	InvalidIdError Error = "The ID you provided is not valid. Please check and try again!"
	ValidationError Error = "Input does not meet requirements. Please check and try again!"

)