
import { afterEach, describe, expect, it, vi } from 'vitest'
import { cleanup, fireEvent, render, screen, waitFor } from '@testing-library/react'
import '@testing-library/jest-dom/vitest'
import App from './App'

afterEach(() => {
  cleanup()
  vi.unstubAllGlobals()
  vi.restoreAllMocks()
})

describe('Calculator', () => {
  it('renders the calculator interface', () => {
    render(<App />)

    expect(screen.getByRole('button', { name: /calculate/i })).toBeInTheDocument()
    expect(screen.getByRole('combobox')).toBeInTheDocument()
  })

  it('displays the result after a successful calculation', async () => {
    const mockFetch = vi.fn().mockResolvedValue({
      ok: true,
      json: async () => ({ result: 50 }),
    })

    vi.stubGlobal('fetch', mockFetch)

    render(<App />)

    fireEvent.change(screen.getByRole('combobox'), {
      target: { value: 'multiply' },
    })

    fireEvent.click(screen.getByRole('button', { name: /calculate/i }))

    await waitFor(() => {
      expect(screen.getByText('50')).toBeInTheDocument()
    })

    expect(mockFetch).toHaveBeenCalledWith(
      '/api/calculate',
      expect.objectContaining({
        method: 'POST',
      }),
    )
  })

  it('displays an error when the API rejects the calculation', async () => {
    const mockFetch = vi.fn().mockResolvedValue({
      ok: false,
      json: async () => ({ error: 'Division by zero is not allowed' }),
    })

    vi.stubGlobal('fetch', mockFetch)

    render(<App />)

    fireEvent.change(screen.getByRole('combobox'), {
      target: { value: 'divide' },
    })

    fireEvent.click(screen.getByRole('button', { name: /calculate/i }))

    await waitFor(() => {
      expect(
        screen.getByText('Division by zero is not allowed'),
      ).toBeInTheDocument()
    })
  })

  it('does not send a request when an input is empty', async () => {
    const mockFetch = vi.fn()
    vi.stubGlobal('fetch', mockFetch)

    render(<App />)

    const numberInputs = screen.getAllByRole('spinbutton')
    fireEvent.change(numberInputs[0], { target: { value: '' } })

    fireEvent.click(screen.getByRole('button', { name: /calculate/i }))

    await waitFor(() => {
      expect(mockFetch).not.toHaveBeenCalled()
    })
  })
})
