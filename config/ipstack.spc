connection "ipstack" {
  plugin    = "ipstack"
  
  # ipstack requires an API token for all requests, but offers a free tier. 
  # Sign up on the ipstack website(https://ipstack.com) to get your free 
  # token (it looks like `e0067f483763d6132d549234f8a6de22`) and set it:
  #token  = "YOUR_TOKEN_HERE"

  # ipstack restricts HTTPS requests to subscribers. This plugin uses HTTP by 
  # default for convenience in the free tier. If you wish to use HTTPS please 
  # set the `https` argument to true:
  #https   = "true"
} 