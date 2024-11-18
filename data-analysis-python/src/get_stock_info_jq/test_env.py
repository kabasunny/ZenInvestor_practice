from dotenv import load_dotenv
import os

# .envファイルを読み込む
load_dotenv()

email = os.getenv("JQUANTS_EMAIL")
password = os.getenv("JQUANTS_PASSWORD")

print(f"Email: {email}")
print(f"Password: {password}")
