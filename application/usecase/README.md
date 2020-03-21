## application/usecase 
- システム仕様上のユースケース群
- `handler`からcallされる
- handler : usecase = 1 : 1 となる
- 基本的に `domain/service/*` をcallし、業務フローを組み立てる