package app

import "github.com/revel/revel"

func init() {
	// Filters is the default set of global filters.
	revel.Filters = []revel.Filter{
		revel.PanicFilter,             // Recover from panics and display an error page instead.
		revel.RouterFilter,            // Use the routing table to select the right Action
		revel.FilterConfiguringFilter, // A hook for adding or removing per-Action filters.
		revel.ParamsFilter,            // Parse parameters into Controller.Params.
		revel.SessionFilter,           // Restore and write the session cookie.
		revel.FlashFilter,             // Restore and write the flash cookie.
		revel.ValidationFilter,        // Restore kept validation errors and save new ones from cookie.
		revel.I18nFilter,              // Resolve the requested language
		revel.InterceptorFilter,       // Run interceptors around the action.
		revel.CompressFilter,          // Compress the result.
		revel.ActionInvoker,           // Invoke the action.
		CORSFilter,                    // Cross Origin Resource Sharing 
	}
}

var CORSFilter = func(c *revel.Controller, fc []revel.Filter) {
        c.Response.Out.Header().Set("Access-Control-Allow-Origin", "*")
        c.Response.Out.Header().Set("Access-Control-Allow-Methods", "GET")
	c.Response.Out.Header().Set("Access-Control-Allow-Headers", "*")
        // Stop here for a Preflighted OPTIONS request
        if c.Request.Method == "OPTIONS" {
                return
        }

        fc[0](c, fc[1:]) // Execute the next filter stage.                                                                                                                                                  
}
