import os
import requests
import json

OPENROUTER_API_URL = "https://openrouter.ai/api/v1/chat/completions"
OPENROUTER_API_KEY = os.getenv("DEEPSEEK")

def analyze_customers_with_gemma(customers: list[dict]) -> str:
    user_content = json.dumps({
        "customers": customers,
        "tasks": [
            "segment_customers",
            "identify_top_clients",
            "sales_risk",
            "recommendations"
        ]
    }, ensure_ascii=False)
    user_content = user_content.encode('utf-8').decode('utf-8')

    payload = {
        "model": "nex-agi/deepseek-v3.1-nex-n1:free",
        "messages": [
            {"role": "system", "content": "You are a CRM analytics assistant. Return strict JSON only."},
            {"role": "user", "content": user_content}
        ],
        "temperature": 0.2,
        "max_tokens": 1000
    }

    headers = {
        "Authorization": f"Bearer {OPENROUTER_API_KEY}",
        "Content-Type": "application/json"
    }

    try:
        response = requests.post(OPENROUTER_API_URL, headers=headers, json=payload, timeout=60)
        response.raise_for_status()
        result = response.json()["choices"][0]["message"]["content"]
    except (requests.exceptions.HTTPError, KeyError, IndexError) as e:
        print(f"OpenRouter API error: {e}")
        result = "{}"
    return result
