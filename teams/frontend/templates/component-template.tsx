/**
 * ComponentName
 *
 * Brief description of what this component does.
 *
 * @param props - Component properties
 * @returns JSX element
 */

import { useState } from 'react'

// ============================================================================
// Types
// ============================================================================

interface ComponentNameProps {
  /** Description of prop1 */
  prop1: string
  /** Description of prop2 */
  prop2?: number
  /** Optional callback function */
  onAction?: (value: string) => void
}

// ============================================================================
// Component
// ============================================================================

/**
 * ComponentName component
 */
export function ComponentName({
  prop1,
  prop2 = 0,
  onAction
}: ComponentNameProps) {
  // ------------------------------------------------------------------------
  // State
  // ------------------------------------------------------------------------
  const [localState, setLocalState] = useState<string>('')

  // ------------------------------------------------------------------------
  // Derived State
  // ------------------------------------------------------------------------
  const derivedValue = `${prop1}-${prop2}`

  // ------------------------------------------------------------------------
  // Effects
  // ------------------------------------------------------------------------
  // useEffect(() => {
  //   // Effect logic here
  // }, [dependencies])

  // ------------------------------------------------------------------------
  // Handlers
  // ------------------------------------------------------------------------
  const handleClick = () => {
    onAction?.(localState)
  }

  // ------------------------------------------------------------------------
  // Render
  // ------------------------------------------------------------------------
  return (
    <div className="component-name">
      <h2>{prop1}</h2>
      <p>{derivedValue}</p>
      <button onClick={handleClick}>
        Action
      </button>
    </div>
  )
}

// ============================================================================
// Default Export
// ============================================================================

export default ComponentName
