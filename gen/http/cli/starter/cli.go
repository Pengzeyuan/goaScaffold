// Code generated by goa v3.1.1, DO NOT EDIT.
//
// starter HTTP client CLI support package
//
// Command:
// $ goa gen starter/design

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	userc "starter/gen/http/user/client"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `user (login-by-username|login-by-sms-code|update-password|get-captcha-image|send-sms-code)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` user login-by-username --body '{
      "captchaId": "2qb",
      "humanCode": "7wt",
      "password": "password",
      "username": "user"
   }'` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (goa.Endpoint, interface{}, error) {
	var (
		userFlags = flag.NewFlagSet("user", flag.ContinueOnError)

		userLoginByUsernameFlags    = flag.NewFlagSet("login-by-username", flag.ExitOnError)
		userLoginByUsernameBodyFlag = userLoginByUsernameFlags.String("body", "REQUIRED", "")

		userLoginBySmsCodeFlags    = flag.NewFlagSet("login-by-sms-code", flag.ExitOnError)
		userLoginBySmsCodeBodyFlag = userLoginBySmsCodeFlags.String("body", "REQUIRED", "")

		userUpdatePasswordFlags     = flag.NewFlagSet("update-password", flag.ExitOnError)
		userUpdatePasswordBodyFlag  = userUpdatePasswordFlags.String("body", "REQUIRED", "")
		userUpdatePasswordTokenFlag = userUpdatePasswordFlags.String("token", "REQUIRED", "")

		userGetCaptchaImageFlags = flag.NewFlagSet("get-captcha-image", flag.ExitOnError)

		userSendSmsCodeFlags    = flag.NewFlagSet("send-sms-code", flag.ExitOnError)
		userSendSmsCodeBodyFlag = userSendSmsCodeFlags.String("body", "REQUIRED", "")
	)
	userFlags.Usage = userUsage
	userLoginByUsernameFlags.Usage = userLoginByUsernameUsage
	userLoginBySmsCodeFlags.Usage = userLoginBySmsCodeUsage
	userUpdatePasswordFlags.Usage = userUpdatePasswordUsage
	userGetCaptchaImageFlags.Usage = userGetCaptchaImageUsage
	userSendSmsCodeFlags.Usage = userSendSmsCodeUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "user":
			svcf = userFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "user":
			switch epn {
			case "login-by-username":
				epf = userLoginByUsernameFlags

			case "login-by-sms-code":
				epf = userLoginBySmsCodeFlags

			case "update-password":
				epf = userUpdatePasswordFlags

			case "get-captcha-image":
				epf = userGetCaptchaImageFlags

			case "send-sms-code":
				epf = userSendSmsCodeFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "user":
			c := userc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "login-by-username":
				endpoint = c.LoginByUsername()
				data, err = userc.BuildLoginByUsernamePayload(*userLoginByUsernameBodyFlag)
			case "login-by-sms-code":
				endpoint = c.LoginBySmsCode()
				data, err = userc.BuildLoginBySmsCodePayload(*userLoginBySmsCodeBodyFlag)
			case "update-password":
				endpoint = c.UpdatePassword()
				data, err = userc.BuildUpdatePasswordPayload(*userUpdatePasswordBodyFlag, *userUpdatePasswordTokenFlag)
			case "get-captcha-image":
				endpoint = c.GetCaptchaImage()
				data = nil
			case "send-sms-code":
				endpoint = c.SendSmsCode()
				data, err = userc.BuildSendSmsCodePayload(*userSendSmsCodeBodyFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// userUsage displays the usage of the user command and its subcommands.
func userUsage() {
	fmt.Fprintf(os.Stderr, `微服务
Usage:
    %s [globalflags] user COMMAND [flags]

COMMAND:
    login-by-username: 使用账号密码登录
    login-by-sms-code: 使用短信验证码登录
    update-password: 修改登录密码
    get-captcha-image: 获取图形验证码
    send-sms-code: 发送短信验证码

Additional help:
    %s user COMMAND --help
`, os.Args[0], os.Args[0])
}
func userLoginByUsernameUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] user login-by-username -body JSON

使用账号密码登录
    -body JSON: 

Example:
    `+os.Args[0]+` user login-by-username --body '{
      "captchaId": "2qb",
      "humanCode": "7wt",
      "password": "password",
      "username": "user"
   }'
`, os.Args[0])
}

func userLoginBySmsCodeUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] user login-by-sms-code -body JSON

使用短信验证码登录
    -body JSON: 

Example:
    `+os.Args[0]+` user login-by-sms-code --body '{
      "mobile": "guest",
      "verify_code": "123456"
   }'
`, os.Args[0])
}

func userUpdatePasswordUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] user update-password -body JSON -token STRING

修改登录密码
    -body JSON: 
    -token STRING: 

Example:
    `+os.Args[0]+` user update-password --body '{
      "new_password": "abc123",
      "old_password": "123abc"
   }' --token "eyJhbGciOiJIUz..."
`, os.Args[0])
}

func userGetCaptchaImageUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] user get-captcha-image

获取图形验证码

Example:
    `+os.Args[0]+` user get-captcha-image
`, os.Args[0])
}

func userSendSmsCodeUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] user send-sms-code -body JSON

发送短信验证码
    -body JSON: 

Example:
    `+os.Args[0]+` user send-sms-code --body '{
      "captchaId": "4v",
      "humanCode": "ir6",
      "mobile": "7qw"
   }'
`, os.Args[0])
}
