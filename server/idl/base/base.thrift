namespace go base

// 错误码枚举定义
enum ErrCode {
    // 成功
    Success = 0
    
    // 通用错误 1xxxx
    NoRoute = 10001
    NoMethod = 10002
    BadRequest = 10003
    ParamsErr = 10004
    AuthorizeFail = 10005
    TooManyRequest = 10006
    ServiceErr = 10007
    RecordNotFound = 10008
    RecordAlreadyExist = 10009
    
    // 用户服务错误 2xxxx
    RPCUserSrvErr = 20001
    UserSrvErr = 20002
    UserNotFound = 20003
    PasswordError = 20004
    EmailAlreadyExist = 20005
    InvalidPassword = 20006
    PhoneAlreadyExist = 20007
    UsernameAlreadyExist = 20008
    VerifyCodeError = 20009
    VerifyCodeExpired = 20010
    VerifyCodeSendFailed = 20011
    TooManyVerifyCodeRequest = 20012
    
    // 面试服务错误 3xxxx
    RPCInterviewSrvErr = 30001
    InterviewSrvErr = 30002
    InterviewNotFound = 30003
    InterviewAlreadyFinished = 30004
    InvalidInterviewType = 30005
    
    // 题库服务错误 4xxxx
    RPCQuestionSrvErr = 40001
    QuestionSrvErr = 40002
    QuestionNotFound = 40003
    InvalidDifficulty = 40004
    CategoryNotFound = 40005
    
    // 存储服务错误 5xxxx
    RPCStorageSrvErr = 50001
    StorageSrvErr = 50002
    FileUploadError = 50003
    FileNotFound = 50004
    InvalidFileType = 50005
    FileSizeExceeded = 50006
}