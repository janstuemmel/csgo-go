## Log server preperation

A list of useful cvars to prepare a server for logging:

```sh
log on|off              # Enable/Disable logging

sv_logsecret  0         # =1 Activates logsecret, set a custom secret with  logaddress_token_secret

mp_logdetail         0          # Logs attacks. Values are: 0=off, 1=enemy, 2=teammate, 3=both)
mp_logdetail_items   0          # Logs a line any time a player acquires or loses an item.
mp_logmoney          0          # Enables money logging. Values are: 0=off, 1=on 
mp_logdistance_2d 	 250        # Enables distance logging every so many units
mp_logdistance_sec 	 15 	      # Enables distance logging every so many seconds
mp_logloadouts 	     1   	      # Enables distance logging with full loadouts 

logaddress_add            <host:port>   # Adds a logaddress
logaddress_add_ex         <host:port>   # Adds a logaddress and generates a secret
logaddress_token_secret   <string>      # Sets a token secret
logaddress_add_ts         <host:port>   # Adds a logaddress and uses the token_secret from above
logaddress_del            <host:port>   # Deletes specific logaddress 
logaddress_delall                       # Deletes all logaddresses
logaddress_list                         # Lists all udp logaddresses
```

If you want to run a csgo game (e.g. with bots) in the background, make sure `mp_autokick 0`, then you won't get kicked when afk. 


## Resources

* [CSGO Cvars](https://developer.valvesoftware.com/wiki/List_of_CS:GO_Cvars)