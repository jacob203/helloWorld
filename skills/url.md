# put version in the url NOT in query part
Specifying the version ensures that the service does not return response elements that your application is not designed to handle.

if you put the version in the query, clients have to handle the version response. 
the old version should not be supported.  
and  there should be a no version url, which means that it is the latest one.

# no need to do retries when 4xx is returned.

# actions
idempotent: return the same result no matter how many times it is called.
* get: get a resource, idempotent
* post: create a resource, it is not idempotent
* put: update a resource, it is **idempotent**
* delete: delete a resource, it is idempotent
* patch: update a resource partly, it is idempotent
