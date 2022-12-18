# 저장소 설명
`백엔드를 위한 GO 프로그래밍` 실습 저장소

# 메모
### 빌드 
빌드 커맨드
```
go mod init <<module name>>
go build
```

### 실행
빌드를 하면 컴파일 결과로 실행 파일이 만들어진다
```
예: 파일 이름이 helloworld라면, 아래 커맨드로 실행한다
./helloworld
```

### 변수
변수 선언
```
var name = "test name"
```

상수 선언
```
const name = "const name"
```

대소문자 구분
- 대문자: public
- 소문자: private

배열 선언
```
// 길이가 고정되어 있지 않은 배열 선언
names := []string{"test1", "test2", "test3"}

// 항목 추가
names = append(names, "test4")

// 초기 메모리 공간이 사이즈의 배열 선언
// make 메서드 사용
names := make([]string, 초기 사이즈)
```

### 함수
기본 구조
```
func <<function name>>(<<variable name>> <<variable type>>) <return type> { }
```

defer 함수
- 메서드가 실행되고 반환되기 직전에 호출되는 함수
- 사용방법
   - 기존 함수 호출: `defer <<function call>>`
   - 인라인 함수: 
      ```go
      defer func() {
        // content
      }()
      ```
### 구조체
go는 객체지향프로그램이 아니다. 그래서 클래스 개념이 없고, 구조체가 있다. 

구조체 샘플
```go
type User struct {
	Name string
	Age  int
}
```

구조체 값 참조
- 기본적으로는 값의 복사가 이루어 진다
- `*`, `&` 키워드를 이용해서 참조로 변경할 수 있다
   ```go
   // 구현
   func incrementAgeReference(user *User) { }

   // 호출
   incrementAgeReference(&myUser)
   ```

인스턴스 메서드
- 익스텐션 메서드 같다
- 구조체에 기능을 확장하는 함수를 추가하는 기능
- 샘플 코드
   ```go
   // 구현
   func (user User) prettyString() string {
	return fmt.Sprintf("Pretty string %s", user.Name)
   }
   // 호출
   myUser.prettyString()
   ```
