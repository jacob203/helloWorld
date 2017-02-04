#reactor and proactor
The reactor design pattern is an event handling pattern for handling service requests delivered concurrently to a service handler by one or more inputs. The service handler then demultiplexes the incoming requests and dispatches them synchronously to the associated request handlers.
the proactor is the same, the only difference is that it calls the associated request handlers asynchronously.

# How to change web IP map
there is a file /private/etc/hosts, DNS router lookups in the file and then requests ip from DNS server.  when you want your computer not to access a url, you can set that website to 0.0.0.0, just like the following:   
0.0.0.0 www.facebook.com  
then you will not access www.facebook.com.
when those changes are done, close the terminal, then the changes are effective.