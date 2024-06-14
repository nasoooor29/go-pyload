// this file for all accounts related things like auth and info about the acc
package gopyload

import (
	"gopyload/models"
	"net/http"
	"net/url"
)

type Pyload struct {
	Address string
	client  *client
	User    string
	Pass    string
}

func NewPyloadAndLogin(address, user, pass string) (*Pyload, error) {
	p := &Pyload{
		Address: address,
		User:    user,
		Pass:    pass,
		client:  NewClient(),
	}
	_, err := p.Login()
	if err != nil {
		return nil, err
	}
	return p, nil
}
func NewPyload(address, user, pass string) *Pyload {
	return &Pyload{
		Address: address,
		User:    user,
		Pass:    pass,
		client:  NewClient(),
	}
}

func (p *Pyload) AddCookie(key, val string) error {
	address, err := url.Parse(p.Address)
	if err != nil {
		return err
	}
	p.client.Client.Jar.SetCookies(address, []*http.Cookie{
		&http.Cookie{
			Name:  key,
			Value: val,
		},
	})
	return nil
}

func (p *Pyload) Login() (models.LoginModel, error) {
	v := url.Values{
		"do":       {"login"},
		"username": {p.User},
		"password": {p.Pass},
		"submit":   {"Login"},
	}
	l, err := GeneralPostEndpoint[models.LoginModel](p, loginEndpoint, v)
	if err != nil {
		val, err := GeneralPostEndpoint[bool](p, loginEndpoint, v)
		if !val {
			return *new(models.LoginModel), HttpMapToStatusCode[http.StatusUnauthorized]
		}
		if err != nil {
			return *new(models.LoginModel), err
		}

	}
	return l, nil
}

func (p *Pyload) Logout() (bool, error) {
	return GeneralGetEndpoint[bool](p, logoutEndpoint)
}

func (p *Pyload) GetUserData(user, pass string) (models.LoginModel, error) {
	return GeneralPostEndpoint[models.LoginModel](p, loginEndpoint, url.Values{
		"username": {user},
		"password": {pass},
	})
}

func (p *Pyload) IsAuthorized(e endpoint, username, password string) (models.LoginModel, error) {
	return GeneralPostEndpoint[models.LoginModel](p, e, url.Values{
		"func":     {e.GetFuncName()},
		"username": {username},
		"password": {password},
	})
}

func (p *Pyload) GetAllUserData() (models.AllUsersDataModel, error) {
	return GeneralPostEndpoint[models.AllUsersDataModel](p, getAllUserDataEndpoint, nil)
}
