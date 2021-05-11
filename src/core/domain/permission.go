package domain

//User related permissions
var ( //This user can:
	PermCreateUser     = 1 //Create users
	PermEditUser       = 2 //Edit created users
	PermDeleteUser     = 4 //Delete created users
	PermReadAllUser    = 8 //Read an user profile
	PermFullUserAccess = PermCreateUser & PermEditUser & PermDeleteUser & PermReadAllUser
)
var UserPermissionIdentifier = "permissions.usermanagement"
