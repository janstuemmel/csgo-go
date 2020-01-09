## Log message specification

### Header

The header of a csgo log message received via udp can have 4 
different structures depending on differnt cvars. 

1. Log message without secret nor token:

```sh
sv_logsecret 0
logaddress_add <ip:port>

# Log output:
����RL 01/09/2020 - 15:45:35: World triggered "Round_Start"
```

2. Log message with secret but without token

```sh
sv_logsecret 2foo
logaddress_add <ip:port>

# Log output:
����S2fooL 01/09/2020 - 15:45:35: World triggered "Round_Start"
```

3. Log message without secret but with token

```sh
sv_logsecret 0
logaddress_token_secret s3cret # returns `logaddress_token_secret:  token checksum = B8032A3B450FB7A6`
logaddress_add_ts <ip:port>

# Log output:
����RTB8032A3B450FB7A6 L 01/09/2020 - 15:45:35: World triggered "Round_Start"
```

4. Log message witho secret and token

```sh
sv_logsecret 2foo
logaddress_token_secret s3cret # returns `logaddress_token_secret:  token checksum = B8032A3B450FB7A6`
logaddress_add_ts <ip:port>

# Log output:
����S2fooTB8032A3B450FB7A6 L 01/09/2020 - 15:45:35: World triggered "Round_Start"
```

Lets look at the header of example 4: `����S2fooTB8032A3B450FB7A6 `:

* `����`: 4 bytes of `0xff`
* `S2foo`: **R** when no secret *or* **S** when secret is set **followed** by a **string** (a secret always beginns with a positive integer)
* `TB8032A3B450FB7A6 `: none *or* **T** followed by the hash of the token, ends with a space
* `L `: L **followed** by a **space**, begin of log message