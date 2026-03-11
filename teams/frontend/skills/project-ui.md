---
name: ui-theme
description: Expert in using the Electric Court design system for building theme-compliant UI components with MUI v7
---
# Skill: Project UI

## Description

Expert in applying the Electric Court design system when building UI components for the Badminton Admin Panel. Ensures all components use theme tokens, follow mobile-first responsive patterns, and maintain visual consistency.

## When to Use This Skill

- Building any new UI component or page
- Styling existing components
- Creating forms, cards, tables, or layouts
- Implementing responsive designs
- Choosing colors, typography, or spacing

## Core Principles

1. **No Hardcoded Values**: Always use theme tokens
2. **Mobile-First**: Design for mobile, enhance for desktop
3. **No Gradients**: Solid colors only
4. **4px Grid**: All spacing in multiples of 4
5. **Theme Tokens Only**: Access via `theme` or `sx` prop

---

## 1. Color System

### Color Tokens

**Primary (Teal)** - Brand identity, court green

- `primary.main`, `primary.light`, `primary.dark`
- Usage: Buttons, brand elements, CTAs

**Secondary (Gold)** - Actions, highlights

- `secondary.main`, `secondary.light`, `secondary.dark`
- Usage: Action buttons, premium features

**Status Colors** - Court availability

- `success.main` - Available (Green)
- `error.main` - Booked (Red)
- `warning.main` - Maintenance (Amber/Orange)
- `info.main` - Reserved/Info

**Text Colors**

- `text.primary` - High contrast, main content
- `text.secondary` - Muted, supporting info
- `text.disabled` - Disabled state

**Background Colors**

- `background.default` - Page background
- `background.paper` - Card/Paper background

### Color Reference Table

| Token                  | Light Mode | Dark Mode | Usage               |
| ---------------------- | ---------- | --------- | ------------------- |
| `primary.main`       | #00A389    | #00A389   | Brand, CTAs         |
| `secondary.main`     | #FFB800    | #FFB800   | Actions, highlights |
| `success.main`       | #10B981    | #34D399   | Available status    |
| `error.main`         | #DC2626    | #EF4444   | Booked status       |
| `warning.main`       | #F59E0B    | #FBBF24   | Maintenance         |
| `background.default` | #F8FAFC    | #0F172A   | Page bg             |
| `background.paper`   | #FFFFFF    | #1E293B   | Card bg             |
| `text.primary`       | #0F172A    | #F1F5F9   | Main text           |
| `text.secondary`     | #64748B    | #94A3B8   | Muted text          |

---

## 2. Typography System

### Typography Variants

### Typography Usage Guide

| Variant  | Size | Weight | Use Case          | Example            |
| -------- | ---- | ------ | ----------------- | ------------------ |
| h1       | 48px | 700    | Hero text         | "Welcome to Admin" |
| h2       | 36px | 700    | Page titles       | "Court Schedule"   |
| h3       | 28px | 600    | Section titles    | "Today's Bookings" |
| h4       | 24px | 600    | Subsections       | "Payment Info"     |
| h5       | 20px | 700    | Component headers | "Court 1"          |
| h6       | 18px | 700    | Small headers     | "Time Slot"        |
| body1    | 16px | 500    | Main content      | "John Doe"         |
| body2    | 14px | 400    | Secondary text    | "09:00 AM"         |
| caption  | 12px | 600    | Badges/labels     | "BOOKED"           |
| overline | 12px | 500    | Categories        | "STATUS"           |

---

## 3. Spacing System

### Spacing Scale (4px baseline)

**Spacing Function**: `theme.spacing(n)` where n × 4 = pixels

**Guidelines**:

- **Micro-spacing (4-8px)**: Inside components, icon gaps
- **Component spacing (12-16px)**: Between related elements
- **Section spacing (24-32px)**: Between content blocks
- **Layout spacing (40-48px)**: Major page divisions

**Common Values**:

- `p: 1` = 4px
- `p: 2` = 8px
- `p: 3` = 12px
- `p: 4` = 16px (default)
- `p: 6` = 24px
- `p: 8` = 32px
- `p: 12` = 48px

### Spacing Reference

| Value       | Pixels | Usage                            |
| ----------- | ------ | -------------------------------- |
| spacing(1)  | 4px    | Icon gaps, tight padding         |
| spacing(2)  | 8px    | Button padding, list gaps        |
| spacing(3)  | 12px   | Card internal spacing            |
| spacing(4)  | 16px   | Standard card padding, form gaps |
| spacing(6)  | 24px   | Section dividers                 |
| spacing(8)  | 32px   | Page margins                     |
| spacing(10) | 40px   | Hero spacing                     |
| spacing(12) | 48px   | Major sections                   |

---

## 4. Responsive Component Sizing

### Mobile-First Approach

**Breakpoint Detection**:

```tsx
const isMobile = useMediaQuery(theme.breakpoints.down('md'));
```

**Responsive Sizing Pattern**:

```tsx
size={isMobile ? 'large' : 'medium'}
sx={{ p: { xs: 4, md: 6 } }}
```

### Responsive Sizing Reference

| Component    | Mobile (<900px) | Desktop (≥900px) | Dense        |
| ------------ | --------------- | ----------------- | ------------ |
| Button       | 48px (large)    | 40px (medium)     | 32px (small) |
| Input        | 56px            | 48px              | 40px         |
| Card padding | 16px            | 24px              | -            |
| Table row    | 56px            | 40px              | 32px         |
| Icon         | 24px            | 20px              | 16px         |
| Touch target | 48×48px min    | 32×32px min      | -            |

---

## 5. Component Overrides

### Pre-Configured Components

**Buttons**

- No elevation by default
- 12px border radius
- Responsive sizing (large/medium/small)
- No text transform

**Text Fields**

- 8px border radius
- Proper padding (12px 16px)
- Responsive sizing

**Cards**

- No elevation
- 1px border with divider color
- 12px border radius

**Chips**

- Default: 32px height
- Small: 24px height
- Rounded borders

**Tables**

- Dense sizing (40px rows)
- Proper cell padding (8px 12px)
- 13px font size

**Dialogs/Drawers**

- 12px border radius
- Proper padding
- fullScreen on mobile

---

## 6. Layout Constants

### Accessing Layout Values

```tsx
const theme = useTheme();
theme.layout.navbarHeight // "64px"
theme.layout.sidebarWidthExpanded // "260px"
theme.layout.sidebarWidthCollapsed // "80px"
theme.layout.appBarBottomHeight // "72px"
```

### Layout Constants Reference

| Constant                               | Value | Usage                    |
| -------------------------------------- | ----- | ------------------------ |
| `theme.layout.navbarHeight`          | 64px  | Top navbar height        |
| `theme.layout.sidebarWidthExpanded`  | 260px | Desktop sidebar width    |
| `theme.layout.sidebarWidthCollapsed` | 80px  | Collapsed sidebar width  |
| `theme.layout.appBarBottomHeight`    | 72px  | Mobile bottom nav height |

---

## 7. Anti-Patterns (What NOT to Do)

❌ **Hardcoded Colors**: Use `bgcolor: 'primary.main'` not `'#00A389'`
❌ **Hardcoded Spacing**: Use `p: 4` not `padding: '16px'`
❌ **Using Gradients**: Solid colors only
❌ **Ignoring Responsive Sizing**: Use `useMediaQuery` for mobile/desktop
❌ **Wrong Typography Variants**: h5 for component headers, not h1
❌ **Inconsistent Border Radius**: Use theme default (12px)
❌ **Small Touch Targets**: Minimum 48×48px on mobile

---

## 8. Quick Reference

### Common Patterns

**Responsive Padding**: `sx={{ p: { xs: 3, md: 6 } }}`
**Responsive Margin**: `sx={{ m: { xs: 2, md: 4 } }}`
**Flex Container**: `sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}`
**Responsive Display**: `sx={{ display: { xs: 'none', md: 'block' } }}`
**Hover State**: `sx={{ '&:hover': { bgcolor: 'primary.light' } }}`

### Breakpoints

```tsx
const isMobile = useMediaQuery(theme.breakpoints.down('md'));
const isDesktop = useMediaQuery(theme.breakpoints.up('md'));
```

### Color Usage

**Background**: `bgcolor: 'primary.main'`, `bgcolor: 'background.paper'`
**Text**: `color: 'text.primary'`, `color: 'text.secondary'`
**Border**: `borderColor: 'divider'`

---

## Execution Rules

1. **Always use theme tokens** - Never hardcode colors, spacing, or sizing
2. **Mobile-first approach** - Design for mobile, enhance for desktop
3. **Use responsive sizing** - Buttons, inputs, cards adapt to viewport
4. **Follow typography hierarchy** - Use correct variant for each use case
5. **Apply 4px spacing grid** - All spacing in multiples of 4
6. **No gradients** - Solid colors only
7. **Minimum touch targets** - 48×48px on mobile, 32×32px on desktop
8. **Use layout constants** - Access via `theme.layout.*`
9. **Leverage component overrides** - MUI components pre-configured
10. **Test responsive behavior** - Use `useMediaQuery` for breakpoints

## Related Workflows

For detailed component examples and implementation patterns, see:

- `/create-theme-compliant-ui` - Step-by-step guide for building common UI patterns
