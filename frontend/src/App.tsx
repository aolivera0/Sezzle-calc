import React, { useState } from "react";
import "./App.css";

function App() {
  const [a, setA] = useState("10");
  const [b, setB] = useState("5");
  const [operation, setOperation] = useState("add");
  const [result, setResult] = useState<number | null>(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState("");

  async function handleCalculate(event: React.FormEvent<HTMLFormElement>) {
    event.preventDefault();

    setError("");
    setResult(null);

    if (a.trim() === "" || b.trim() === "") {
      setError("Please enter both numbers.");
      return;
    }

    const numberA = Number(a);
    const numberB = Number(b);

    if (!Number.isFinite(numberA) || !Number.isFinite(numberB)) {
      setError("Please eneter valid numbers.");
      return;
    }

    try {
      const response = await fetch("/api/calculate", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          operation,
          a: numberA,
          b: numberB,
        }),
      });

      const data = await response.json();

      if (!response.ok) {
        throw new Error(
          typeof data.error === "string"
            ? data.error
            : typeof data.message === "string"
              ? data.message
              : "The calculation failed.",
        );
      }

      if (typeof data.result !== "number" || !Number.isFinite(data.result)) {
        throw new Error("The server returned an invalid result.");
      }

      setResult(data.result);
    } catch (err) {
      setError(
        err instanceof Error
          ? err.message
          : "Unable to connect to the calculator service.",
      );
    } finally {
      setLoading(false);
    }
  }

  return (
    <main className="calculator-page">
      <header className="page-header">
        <p className="eyebrow">Angel H. Olivera Pinzón</p>
        <h1>
          Sezzle technical assesstment 
        </h1>
        <p className="subtitle">
          A calculator powered by a React interface and Go API.
        </p>
      </header>

      <section className="calculator-card">
        <form onSubmit={handleCalculate}>
          <div className="fields">
            <label htmlFor="">
              First number (a)
              <input
                type="number"
                step="any"
                value={a}
                onChange={(e) => setA(e.target.value)}
                placeholder="Enter a number"
              />
            </label>

            <label>
              Second number (b)
              <input
                type="number"
                step="any"
                value={b}
                onChange={(e) => setB(e.target.value)}
                placeholder="Enter a number"
              />
            </label>
          </div>

          <label className="operation-field">
            Operation
            <select
              value={operation}
              onChange={(e) => setOperation(e.target.value)}
            >
              <option value="add">Addition (a+b)</option>
              <option value="subtract">Subtraction (a-b)</option>
              <option value="multiply">Multiplication (a×b)</option>
              <option value="divide">Division (a/b)</option>
            </select>
          </label>

          <button type="submit" disabled={loading}>
            {loading ? "Calculating..." : "Calculate result"}
          </button>
        </form>

        <div className="feedback-area" aria-live="polite">
          {error ? (
            <p className="error-message" role="alert">
              {error}
            </p>
          ) : result !== null ? (
            <div className="result-panel">
              <p>Result</p>
              <output>{result}</output>
            </div>
          ) : (
            <div className="feedback-placeholder" aria-hidden="true" />
          )}
        </div>

        <p className="card-footer">
          Calculations are processed by the Go backend.
        </p>
      </section>
    </main>
  );
}

export default App;
