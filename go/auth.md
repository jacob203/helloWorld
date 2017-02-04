difference between HMAC and encription
these days I investigated the golang package "securecookie", I was a little confused about auth key and block key, the document says an auth key is required and a block key is optional, the block key is used to encrypt the cookie value.  
so I did not understand what the difference was.
after googling it, I found that there are three concepts I need to know first, authentication, authorization and encryption.
* authentication
it is used to identify who you are. for a server, it needs to know who is accessing its information, generally a client needs to input user name and password, for a client, it needs to know who it is accessing, usually it needs an certification issued by a trusted thirdparty.
* Authorization
it is used for a server to check a client what it can access and can't, it is based on authentication.
* Encryption
it is used to encrypt a message which can only be seen by a client who has a key to decrypt.

# Cryptographic
it has three important principles:
1. Integrity: make sure the message has not been modified
2. Authentication: make sure the message is sent by the sender you want
3. Non-repudiation: 
4. Keys: does the receiver need a shared secret key, or public-private keypairs?
### hash function: MD5, sha1, and sha2
genenally its input is a string, It is a mathematical algorithm that maps data of arbitrary size to a bit string of a fixed size (a hash function) which is designed to also be a **one-way** function. only the same string can get the same hash value, so usually it is used in the following example:
1. verify if the message is changed or not, such as checksum
2. store password, right now it is not safe to store the clear password in db, use hash function to get a hash value to be stored in db, it can be used to verify the password, and it is also safe, even if the db is hacked.
3. use hash value to find a file, like git
4. send a big file: send the file hash value in a secure channel, then send the file in an insecure channel, when the download is finished, use the hash value to verify its integrity.
by the way, right now some methods such as rainbow tables have been published which allow people to reverse MD5 hashes, so generally, we use good salts to make it hard to happen.
### Encryption
it turns data into a series of unreadable characters, that aren't of a fixed length. the key difference between encryption and hashing is that encrypted strings can be reversed back into their original decryptd form if you have the right key.  
There are two primary types of encryption, symmetric key encryption and public key encryption. symmetric key encryption requires the both side have the same key, the encrypted message with one key can only be dencrypted with the same key. public key encryption has two different keys, one used to encrypt the string with public key and the other uses private key to decrypt it.
* AES: it is the gold standard when it comes to symmetric key encryption
* PGP: it is the most popular public key encryption algorithm

if the raw value doesn't need to known for the other side, hash should always be used, as it is more secure. while you are trying to send secure messages to the other, you need to choose encryption, symmertric encryption provides improved performance and is simpler to use, however the key needs to known by both. so many people share an encryption key with someone using public key encryption, then setup a serure channel using th symmetric encryption.
##### RSA
it is a common public key encryption method, it has a pair of keys, one is private key, the other is public key. generally use public key to encrypt data, and use private key to decrypt data.
When encrypting, you use their public key to write message and they use their private key to read it.
When signing, you use your private key to write message's signature, and they use your public key to check if it's really yours, it is a mechanism to identify who you are.
A signature is proof that the signer has the private key that matches some public key. To do this, it would be enough to encrypt the message with that sender's private key, and include the encrypted version alongside the plaintext version. To verify the sender, decrypt the encrypted version, and check that it is the same as the plaintext. so the plaintext should be exchanged too. 
So:
To generate a signature, make a hash from the plaintext, encrypt it with your private key, include it alongside the plaintext.
To verify a signature, make a hash from the plaintext, decrypt the signature with the sender's public key, check that both hashes are the same.

When you generate a key pair, it's completely arbitrary which one you choose to be the public key, and which is the private key. If you encrypt with one, you can decrypt with the other - it works in both directions.


# how to setup ssl channel
1. Server sends a copy of its asymmetric public key.
2. Browser creates a symmetric session key and encrypts it with the server's asymmetric public key. Then sends it to the server.
usually Brower also send the user's identity such as user name and password, so the server can verify its identity and accept it and setup the serure channel.
3. Server decrypts the encrypted session key using its asymmetric private key to get the symmetric session key.
4. Server and Browser now encrypt and decrypt all transmitted data with the symmetric session key. This allows for a secure channel because only the browser and the server know the symmetric session key, and the session key is only used for that session. If the browser was to connect to the same server the next day, a new session key would be created.

# how https works
it is based on SSL.  
1. A browser requests a secure page (usually https://).
2. The web server sends its public key with its certificate.
3. The browser checks that the certificate was issued by a trusted party (usually a trusted root CA), that the certificate is still valid and that the certificate is related to the site contacted.
4. The browser then uses the public key, to encrypt a random symmetric encryption key and sends it to the server with the encrypted URL required as well as other encrypted http data.
5. The web server decrypts the symmetric encryption key using its private key and uses the symmetric key to decrypt the URL and http data.
6. The web server sends back the requested html document and http data encrypted with the symmetric key.
7. The browser decrypts the http data and html document using the symmetric key and displays the information.

# what is the certificate
How do you know that you are dealing with the right person or rather the right web site. Well, someone has taken great length (if they are serious) to ensure that the web site owners are who they claim to be. This someone, you have to implicitly trust: you have his/her certificate loaded in your browser (a root Certificate). A certificate, contains information about the owner of the certificate, like e-mail address, owner's name, certificate usage, duration of validity, resource location or Distinguished Name (DN) which includes the Common Name (CN) (web site address or e-mail address depending of the usage) and the certificate ID of the person who certifies (signs) this information. It contains also the public key and finally a hash to ensure that the certificate has not been tampered with. As you made the choice to trust the person who signs this certificate, therefore you also trust this certificate. This is a certificate trust tree or certificate path. Usually your browser or application has already loaded the root certificate of well known Certification Authorities (CA) or root CA Certificates. The CA maintains a list of all signed certificates as well as a list of revoked certificates. A certificate is insecure until it is signed, as only a signed certificate cannot be modified. You can sign a certificate using itself, it is called a self signed certificate. All root CA certificates are self signed.

Certificate: 
    Data: 
        Version: 3 (0x2) 
        Serial Number: 1 (0x1) 
        Signature Algorithm: md5WithRSAEncryption 
        Issuer: C=FJ, ST=Fiji, L=Suva, O=SOPAC, OU=ICT, CN=SOPAC Root CA/Email=administrator@sopac.org 
        Validity 
            Not Before: Nov 20 05:47:44 2001 GMT 
            Not After : Nov 20 05:47:44 2002 GMT 
        Subject: C=FJ, ST=Fiji, L=Suva, O=SOPAC, OU=ICT, CN=www.sopac.org/Email=administrator@sopac.org 
        Subject Public Key Info: 
            Public Key Algorithm: rsaEncryption  
            RSA Public Key: (1024 bit) 
                Modulus (1024 bit): 
                    00:ba:54:2c:ab:88:74:aa:6b:35:a5:a9:c1:d0:5a: 
                    9b:fb:6b:b5:71:bc:ef:d3:ab:15:cc:5b:75:73:36: 
                    b8:01:d1:59:3f:c1:88:c0:33:91:04:f1:bf:1a:b4: 
                    7a:c8:39:c2:89:1f:87:0f:91:19:81:09:46:0c:86: 
                    08:d8:75:c4:6f:5a:98:4a:f9:f8:f7:38:24:fc:bd: 
                    94:24:37:ab:f1:1c:d8:91:ee:fb:1b:9f:88:ba:25: 
                    da:f6:21:7f:04:32:35:17:3d:36:1c:fb:b7:32:9e: 
                    42:af:77:b6:25:1c:59:69:af:be:00:a1:f8:b0:1a: 
                    6c:14:e2:ae:62:e7:6b:30:e9 
                Exponent: 65537 (0x10001) 
         X509v3 extensions: 
             X509v3 Basic Constraints: 
                 CA:FALSE 
             Netscape Comment: 
                 OpenSSL Generated Certificate
             X509v3 Subject Key Identifier:
                 FE:04:46:ED:A0:15:BE:C1:4B:59:03:F8:2D:0D:ED:2A:E0:ED:F9:2F 
             X509v3 Authority Key Identifier:
                 keyid:E6:12:7C:3D:A1:02:E5:BA:1F:DA:9E:37:BE:E3:45:3E:9B:AE:E5:A6 
                 DirName:/C=FJ/ST=Fiji/L=Suva/O=SOPAC/OU=ICT/CN=SOPAC Root CA/Email=administrator@sopac.org 
                 serial:00
    Signature Algorithm: md5WithRSAEncryption
        34:8d:fb:65:0b:85:5b:e2:44:09:f0:55:31:3b:29:2b:f4:fd: 
        aa:5f:db:b8:11:1a:c6:ab:33:67:59:c1:04:de:34:df:08:57: 
        2e:c6:60:dc:f7:d4:e2:f1:73:97:57:23:50:02:63:fc:78:96: 
        34:b3:ca:c4:1b:c5:4c:c8:16:69:bb:9c:4a:7e:00:19:48:62: 
        e2:51:ab:3a:fa:fd:88:cd:e0:9d:ef:67:50:da:fe:4b:13:c5: 
        0c:8c:fc:ad:6e:b5:ee:40:e3:fd:34:10:9f:ad:34:bd:db:06: 
        ed:09:3d:f2:a6:81:22:63:16:dc:ae:33:0c:70:fd:0a:6c:af:
        bc:5a 
-----BEGIN CERTIFICATE----- 
MIIDoTCCAwqgAwIBAgIBATANBgkqhkiG9w0BAQQFADCBiTELMAkGA1UEBhMCRkox 
DTALBgNVBAgTBEZpamkxDTALBgNVBAcTBFN1dmExDjAMBgNVBAoTBVNPUEFDMQww 
CgYDVQQLEwNJQ1QxFjAUBgNVBAMTDVNPUEFDIFJvb3QgQ0ExJjAkBgkqhkiG9w0B 
CQEWF2FkbWluaXN0cmF0b3JAc29wYWMub3JnMB4XDTAxMTEyMDA1NDc0NFoXDTAy 
MTEyMDA1NDc0NFowgYkxCzAJBgNVBAYTAkZKMQ0wCwYDVQQIEwRGaWppMQ0wCwYD 
VQQHEwRTdXZhMQ4wDAYDVQQKEwVTT1BBQzEMMAoGA1UECxMDSUNUMRYwFAYDVQQD 
Ew13d3cuc29wYWMub3JnMSYwJAYJKoZIhvcNAQkBFhdhZG1pbmlzdHJhdG9yQHNv 
cGFjLm9yZzCBnzANBgkqhkiG9w0BAQEFAAOBjQAwgYkCgYEAulQsq4h0qms1panB 
0Fqb+2u1cbzv06sVzFt1cza4AdFZP8GIwDORBPG/GrR6yDnCiR+HD5EZgQlGDIYI 
2HXEb1qYSvn49zgk/L2UJDer8RzYke77G5+IuiXa9iF/BDI1Fz02HPu3Mp5Cr3e2 
JRxZaa++AKH4sBpsFOKuYudrMOkCAwEAAaOCARUwggERMAkGA1UdEwQCMAAwLAYJ 
YIZIAYb4QgENBB8WHU9wZW5TU0wgR2VuZXJhdGVkIENlcnRpZmljYXRlMB0GA1Ud
DgQWBBT+BEbtoBW+wUtZA/gtDe0q4O35LzCBtgYDVR0jBIGuMIGrgBTmEnw9oQLl 
uh/anje+40U+m67lpqGBj6SBjDCBiTELMAkGA1UEBhMCRkoxDTALBgNVBAgTBEZp 
amkxDTALBgNVBAcTBFN1dmExDjAMBgNVBAoTBVNPUEFDMQwwCgYDVQQLEwNJQ1Qx 
FjAUBgNVBAMTDVNPUEFDIFJvb3QgQ0ExJjAkBgkqhkiG9w0BCQEWF2FkbWluaXN0 
cmF0b3JAc29wYWMub3JnggEAMA0GCSqGSIb3DQEBBAUAA4GBADSN+2ULhVviRAnw 
VTE7KSv0/apf27gRGsarM2dZwQTeNN8IVy7GYNz31OLxc5dXI1ACY/x4ljSzysQb 
xUzIFmm7nEp+ABlIYuJRqzr6/YjN4J3vZ1Da/ksTxQyM/K1ute5A4/00EJ+tNL3b 
Bu0JPfKmgSJjFtyuMwxw/Qpsr7xa
-----END CERTIFICATE-----
As You may have noticed, the certificate contains the reference to the issuer, the public key of the owner of this certificate, the dates of validity of this certificate and the signature of the certificate to ensure this certificate hasen't been tampered with. The certificate does not contain the private key as it should never be transmitted in any form whatsoever. This certificate has all the elements to send an encrypted message to the owner (using the public key) or to verify a message signed by the author of this certificate.

signature means it hashes the data, then add it to the end, in the same form as checksum, so the client with private key can not only decode the message and know whether if it is correct or not.

# summary
RSA can only decript the encripted message, but how to check the decripted message is correct has to use hash to make sure if it is the correct one.



