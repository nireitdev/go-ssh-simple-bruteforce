# go-ssh-simple-bruteforce
Simple POC for a Brute-Force SSH dictorionary password attack.

FOR TESTING PURPOSE AND SELF LEARNING!! 

It's necessary to obtain an already downloaded dictionary with passwords. Typically, "rockyou.txt" (Google it!)

Option available:

-f file text to use as a dictionary. 

-h remote server

-p remote ssh port

-u remote ssh user 

-t number max of threads to create

Recommendation: Use a low number of threads (for example, 4). Many SSH servers drop connections.
For instance, OpenSSH uses the parameter MaxStartups to limit simultaneous connections.

View/Edit runme.sh for a fast overview.



