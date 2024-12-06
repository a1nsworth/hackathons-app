export class BaseApiUrl {
    static readonly Base = 'http://localhost:4242';
}

export class AuthUrl {
    static readonly Base = `${BaseApiUrl.Base}/auth`;
    static readonly Login = `${AuthUrl.Base}/login`;
    static readonly Register = `${AuthUrl.Base}/register`;
    static readonly Refresh = `${AuthUrl.Base}/refresh`;

}

export class HackathonApiUrl {
    static readonly GetAll = `${BaseApiUrl.Base}/hackathons/`;
}