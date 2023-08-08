package gopresence

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"runtime"
	"time"

	np "gopkg.in/natefinch/npipe.v2"
)

type Client struct {
	Conn net.Conn
}

func (c *Client) Connect(clientId string) error {
	var socket net.Conn
	var err error

	o := runtime.GOOS

	if o == "windows" {
		socket, err = np.DialTimeout(`\\.\pipe\discord-ipc-0`, time.Second*2)

		if err != nil {
			return err
		}

	} else if o == "darwin" || o == "linux" {
		paths := []string{"XDG_RUNTIME_DIR", "TMPDIR", "TMP", "TEMP", "/tmp"}

		for _, v := range paths {
			if path, exists := os.LookupEnv(v); exists {
				socket, err = net.DialTimeout("unix", path+`\discord-ipc-0`, time.Second*2)

				if err != nil {
					return err
				}

			}
		}
	} else {
		return fmt.Errorf("%s os is not supported", o)
	}

	c.Conn = socket
	c.Send(0, Handshake{V: 1, ClientID: clientId})
	return nil
}

func (c *Client) Send(op int, data interface{}) error {
	var err error

	payload, _ := json.Marshal(data)

	buffer := new(bytes.Buffer)

	if err = binary.Write(buffer, binary.LittleEndian, int32(op)); err != nil {
		return err
	}

	if err = binary.Write(buffer, binary.LittleEndian, int32(len(payload))); err != nil {
		return err
	}

	if err = binary.Write(buffer, binary.LittleEndian, payload); err != nil {
		return err
	}

	if _, err = c.Conn.Write(buffer.Bytes()); err != nil {
		return err
	}

	return nil
}

func (c *Client) Receive() (string, error) {
	buf := make([]byte, 512)
	_, err := c.Conn.Read(buf)

	if err != nil {
		return "", err
	}

	buffer := new(bytes.Buffer)

	for _, b := range buf[8:] {
		buffer.WriteByte(b)
	}

	r := buffer.String()

	return r, nil
}

func (c *Client) SetActivity(activity Activity) error {
	err := c.Send(1, Payload{
		Cmd:   "SET_ACTIVITY",
		Nonce: fmt.Sprint(time.Now().Unix()),
		Args: SetActivityArg{
			Pid:      os.Getpid(),
			Activity: activity,
		},
	})

	if err != nil {
		return err
	}

	res, err := c.Receive()

	if err != nil {
		return err
	}

	var ret map[string]interface{}

	json.Unmarshal(bytes.Trim([]byte(res), "\x00"), &ret)

	if data, ok := ret["data"]; ok {
		d, ok := data.(map[string]interface{})

		if ok {
			code, ok := d["code"]

			if ok && code.(float64) >= 4000 {
				return fmt.Errorf("code: %v, %s", code, d["message"].(string))
			}
		}
	}

	return nil
}
