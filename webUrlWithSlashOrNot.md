recently I met an issue when I migrated other endpoints, 
but I forgot the slash issue that means some endpoints have trailing slash, while the others doesn't.
and our server enables strict feature, it means both doesn't return 200, only one returns 200 and the other returns 301 or 302.

I googled it, found a page [to slash or not to slash](https://webmasters.googleblog.com/2010/04/to-slash-or-not-to-slash.html)
it says it is a good behavior, it is good for search engines, if you both return 200, for search engines, 
they are two urls although most search engines will eliminate one, it adds extra efforts,
and google index generally stores the url with a trailing slash.

so there is another question: when does url use a trailing slash or not?  
from the meaning of REST, an url is a resource, and an url with a trailing slash means a collection of resources, an url without a trailing slash means a specific resource.  

|   URL    |    Get     |   Put    |    Post   |    Delete   |
|:--------:|:----------:|:--------:|:---------:|:-----------:|
|http://example.com/dogs/|get the dogs list|update the dogs list|create a dog in the dogs collection|delete the entire collection|
|http://example.com/dogs/herry|get the dog information whose name is herry|update the dog herry information|**not generally used**|delete the dog herry from the dogs collection|


