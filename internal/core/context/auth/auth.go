package auth

type Context struct {
	Provider string      // How we authorize
	Subject  interface{} // What do we authorize

	Claims interface{} // Extracted data from authorization
	Result bool        // Is authorization successfull
	Roles  []string    // Roles list from authorization
}
