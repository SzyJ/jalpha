# Japanese-Alphabet
hiragana and katakana tables displayed in your terminal!

## Try it out!
Jalpha is currently hosted on ```api.szy.wtf/jalpha```

Try the following commands in your terminal now!
```bash
curl api.szy.wtf/jalpha/hira
```
```bash
curl api.szy.wtf/jalpha/kata
```
```bash
curl api.szy.wtf/jalpha/hira/s/zu
```
```bash
curl api.szy.wtf/jalpha/kata/s/tsu
```
If Japanese characters are not being displayed in your terminal, you may be using a font that does not support these charcaters.

## Usage
In these examples, it is assumed that the server has been started on localhost on port 8080

To print hiragana table: ```curl localhost:8080/hira```

To print katakana table:```curl localhost:8080/kata```

To search for a certain character with romanji, the previous commands can be appended with:```/s/<Romanji>```

For example, to see the 'shi' hiragana character, ```curl localhost:8080/hira/s/shi``` will return ```shi: し```

If multiple characters fit the search criteria, they are both returned: ```ji: ジ or ヂ```

## Building
### Dependencies
This progam uses [gorilla/mux](https://github.com/gorilla/mux).

### Build and run
Build and run the server:
```
go build && ./jalpha
```

By default, the server will start on port 8080, this can be changed with an argument like so
```
./jalpha <Your_Port_Here>
```

## Future Additions
* Expand alphabet to include Digraphs (e.g. kya: きゃ)
* Extend search feature to replace chars in an arbitrary string to japanese alphabet (e.g. shumon -> シュモン)
