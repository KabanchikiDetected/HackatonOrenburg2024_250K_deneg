FROM python:3.11 as builder


ENV PATH="/root/.local/bin:${PATH}"

WORKDIR /app

RUN curl -sSL https://install.python-poetry.org | python3 -

COPY . .
RUN poetry install --no-dev

CMD ["poetry", "run", "uvicorn", "main:start"]