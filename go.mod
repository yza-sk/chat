module example.com/chat

go 1.23.4

replace example.com/chat/common/message => ../common/message

require github.com/garyburd/redigo v1.6.4
