# Playwright

## Script

```bash
cd playwright/pyenv/src

// open browser & select file
pytest --headed test_example.py
```

### 套件相關

```bash
// 重設依賴項
pip freeze > requirements.txt

// 更新playwright
pip install pytest-playwright playwright -U
```

### Other script

```bash
// headless
pytest
```
