---
description: Guidelines for defining UI content when creating features - titles, sizing, cards, and actions for hybrid mobile-first design
---
# styled component guide for Features

This workflow helps you make consistent decisions about UI content structure when building new features for the Badminton Admin Panel.

## Decision Framework

When creating a new feature, follow this hierarchy to determine proper UI elements:

---

## 1. Page Title Selection

### Title Hierarchy Rules

| Page Type            | Title Variant | Size | Example                                   |
| -------------------- | ------------- | ---- | ----------------------------------------- |
| Dashboard/Home       | `h1`        | 48px | "Welcome to Badminton Admin"              |
| Main Feature Page    | `h2`        | 36px | "Court Bookings", "Customer Management"   |
| Sub-page/Detail      | `h3`        | 28px | "Booking Details", "Court Schedule"       |
| Section within Page  | `h4`        | 24px | "Payment Information", "Customer History" |
| Card/Component Title | `h5`        | 20px | "Court 1", "Recent Bookings"              |
| List Item Title      | `h6`        | 18px | "Time Slot", "Equipment Rental"           |

### Decision Tree

```
Is this the main landing page?
├─ YES → Use h1 (48px) - "Welcome to [App Name]"
└─ NO → Is this a top-level feature?
    ├─ YES → Use h2 (36px) - "[Feature Name]"
    └─ NO → Is this a detail/sub-page?
        ├─ YES → Use h3 (28px) - "[Detail Name]"
        └─ NO → Is this a section within a page?
            ├─ YES → Use h4 (24px) - "[Section Name]"
            └─ NO → Is this a card/component?
                ├─ YES → Use h5 (20px) - "[Component Name]"
                └─ NO → Use h6 (18px) - "[Small Header]"
```

### Examples by Feature

**Bookings Feature**

- Page: `h2` - "Court Bookings"
- Card: `h5` - "Court 1"
- Section: `h4` - "Today's Schedule"
- List item: `h6` - "09:00 AM - 10:00 AM"

**Customer Management**

- Page: `h2` - "Customers"
- Detail page: `h3` - "Customer Details"
- Section: `h4` - "Booking History"
- Card: `h5` - "John Doe"

---

## 2. Content Sizing (Mobile-First)

### Text Content

| Content Type    | Mobile              | Desktop             | Usage              |
| --------------- | ------------------- | ------------------- | ------------------ |
| Main heading    | `h2` (36px)       | `h2` (36px)       | Page title         |
| Section heading | `h4` (24px)       | `h3` (28px)       | Major sections     |
| Card title      | `h5` (20px)       | `h5` (20px)       | Card headers       |
| Body text       | `body1` (16px)    | `body1` (16px)    | Main content       |
| Supporting text | `body2` (14px)    | `body2` (14px)    | Secondary info     |
| Labels          | `overline` (12px) | `overline` (12px) | Field labels       |
| Metadata        | `caption` (12px)  | `caption` (12px)  | Timestamps, badges |

### Component Sizing

| Component    | Mobile (<900px)  | Desktop (≥900px) | Code                                         |
| ------------ | ---------------- | ----------------- | -------------------------------------------- |
| Button       | `large` (48px) | `medium` (40px) | `size={isMobile ? 'large' : 'medium'}`     |
| Input        | 56px height      | 48px height       | `size={isMobile ? 'medium' : 'small'}`     |
| Card padding | 16px (`p: 4`)  | 24px (`p: 6`)   | `p: { xs: 4, md: 6 }`                      |
| Icon         | 24px             | 20px              | `fontSize={isMobile ? 'medium' : 'small'}` |
| Chip         | `small` (24px) | `small` (24px)  | `size="small"`                             |

### Spacing Between Elements

| Context         | Mobile                 | Desktop                | Code                           |
| --------------- | ---------------------- | ---------------------- | ------------------------------ |
| Form fields     | 12px (`spacing={3}`) | 12px (`spacing={3}`) | `<Stack spacing={3}>`        |
| Card sections   | 12px (`spacing={3}`) | 16px (`spacing={4}`) | `spacing={{ xs: 3, md: 4 }}` |
| Page sections   | 16px (`spacing={4}`) | 24px (`spacing={6}`) | `spacing={{ xs: 4, md: 6 }}` |
| Major divisions | 24px (`spacing={6}`) | 32px (`spacing={8}`) | `spacing={{ xs: 6, md: 8 }}` |

---

## 3. Card Structure

### When to Use Cards

✅ **Use cards for:**

- Booking information
- Customer profiles
- Court details
- Summary statistics
- List items on mobile
- Grouped related information

❌ **Don't use cards for:**

- Full-page forms (use Box instead)
- Single text elements
- Navigation items
- Table rows on desktop

### Card Anatomy

```tsx
<Card sx={{ p: { xs: 4, md: 6 } }}>
  <Stack spacing={3}>
    {/* Header Section */}
    <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
      <Typography variant="h5">[Card Title]</Typography>
      <Chip label="[Status]" color="[color]" size="small" />
    </Box>
  
    {/* Content Sections */}
    <Box>
      <Typography variant="overline" color="text.secondary">[Label]</Typography>
      <Typography variant="body1">[Value]</Typography>
    </Box>
  
    <Box>
      <Typography variant="overline" color="text.secondary">[Label]</Typography>
      <Typography variant="body2" color="text.secondary">[Value]</Typography>
    </Box>
  
    {/* Actions */}
    <Stack direction="row" spacing={2}>
      <Button variant="contained" fullWidth>[Primary Action]</Button>
      <Button variant="outlined" fullWidth>[Secondary Action]</Button>
    </Stack>
  </Stack>
</Card>
```

### Card Sizing Guidelines

| Card Type | Mobile Padding  | Desktop Padding | Content Spacing |
| --------- | --------------- | --------------- | --------------- |
| Compact   | `p: 3` (12px) | `p: 4` (16px) | `spacing={2}` |
| Standard  | `p: 4` (16px) | `p: 6` (24px) | `spacing={3}` |
| Spacious  | `p: 5` (20px) | `p: 8` (32px) | `spacing={4}` |

**Default**: Use Standard for most cases

---

## 4. Action Buttons

### Button Placement

**Mobile (<900px)**

- Primary actions: Bottom of screen or card
- Full-width buttons for major actions
- Stack vertically with 8px gap (`spacing={2}`)

**Desktop (≥900px)**

- Primary actions: Top-right or bottom-right
- Inline buttons (not full-width)
- Horizontal layout with 8px gap (`spacing={2}`)

### Button Hierarchy

```tsx
// Primary action (most important)
<Button variant="contained" color="primary" size={isMobile ? 'large' : 'medium'}>
  Confirm Booking
</Button>

// Secondary action (alternative)
<Button variant="outlined" color="primary" size={isMobile ? 'large' : 'medium'}>
  View Details
</Button>

// Tertiary action (less important)
<Button variant="text" size={isMobile ? 'large' : 'medium'}>
  Cancel
</Button>

// Destructive action
<Button variant="contained" color="error" size={isMobile ? 'large' : 'medium'}>
  Delete Booking
</Button>
```

### Action Layout Patterns

**Mobile: Vertical Stack**

```tsx
<Stack spacing={2}>
  <Button variant="contained" fullWidth>Primary Action</Button>
  <Button variant="outlined" fullWidth>Secondary Action</Button>
  <Button variant="text" fullWidth>Cancel</Button>
</Stack>
```

**Desktop: Horizontal Row**

```tsx
<Stack direction="row" spacing={2} justifyContent="flex-end">
  <Button variant="text">Cancel</Button>
  <Button variant="outlined">Secondary Action</Button>
  <Button variant="contained">Primary Action</Button>
</Stack>
```

**Responsive (Recommended)**

```tsx
<Stack 
  direction={{ xs: 'column', md: 'row' }} 
  spacing={2}
  justifyContent={{ md: 'flex-end' }}
>
  <Button variant="text" fullWidth={{ xs: true, md: false }}>Cancel</Button>
  <Button variant="outlined" fullWidth={{ xs: true, md: false }}>View Details</Button>
  <Button variant="contained" fullWidth={{ xs: true, md: false }}>Confirm</Button>
</Stack>
```

---

## 5. Feature Layout Templates

### Template 1: List/Grid View

**Use for**: Bookings list, Customers list, Courts list

```tsx
function FeatureListPage() {
  const theme = useTheme();
  const isMobile = useMediaQuery(theme.breakpoints.down('md'));
  
  return (
    <Box sx={{ p: { xs: 3, md: 6 } }}>
      <Stack spacing={{ xs: 4, md: 6 }}>
        {/* Page Header */}
        <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
          <Typography variant="h2">[Feature Name]</Typography>
          <Button 
            variant="contained" 
            size={isMobile ? 'large' : 'medium'}
            startIcon={<AddIcon />}
          >
            New [Item]
          </Button>
        </Box>
      
        {/* Filters (Optional) */}
        <Card sx={{ p: { xs: 3, md: 4 } }}>
          <Stack spacing={2}>
            <Typography variant="h6">Filters</Typography>
            <Grid container spacing={2}>
              <Grid item xs={12} md={4}>
                <TextField fullWidth label="Search" size={isMobile ? 'medium' : 'small'} />
              </Grid>
              <Grid item xs={12} md={4}>
                <TextField fullWidth label="Status" select size={isMobile ? 'medium' : 'small'}>
                  <MenuItem value="all">All</MenuItem>
                </TextField>
              </Grid>
            </Grid>
          </Stack>
        </Card>
      
        {/* Content: Cards on mobile, Table on desktop */}
        {isMobile ? (
          <Stack spacing={2}>
            {items.map(item => <ItemCard key={item.id} item={item} />)}
          </Stack>
        ) : (
          <TableContainer component={Paper}>
            <Table size="small">
              {/* Table content */}
            </Table>
          </TableContainer>
        )}
      </Stack>
    </Box>
  );
}
```

### Template 2: Detail View

**Use for**: Booking details, Customer profile, Court information

```tsx
function FeatureDetailPage() {
  const theme = useTheme();
  const isMobile = useMediaQuery(theme.breakpoints.down('md'));
  
  return (
    <Box sx={{ p: { xs: 3, md: 6 } }}>
      <Stack spacing={{ xs: 4, md: 6 }}>
        {/* Breadcrumb/Back Navigation */}
        <Button startIcon={<ArrowBackIcon />} onClick={goBack}>
          Back to [List]
        </Button>
      
        {/* Page Header with Actions */}
        <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
          <Typography variant="h3">[Item Name]</Typography>
          <Stack direction="row" spacing={2}>
            <Button variant="outlined" size={isMobile ? 'large' : 'medium'}>Edit</Button>
            <Button variant="contained" color="error" size={isMobile ? 'large' : 'medium'}>
              Delete
            </Button>
          </Stack>
        </Box>
      
        {/* Content Sections */}
        <Grid container spacing={{ xs: 2, md: 3 }}>
          <Grid item xs={12} md={8}>
            <Stack spacing={3}>
              <Card sx={{ p: { xs: 4, md: 6 } }}>
                <Typography variant="h4" gutterBottom>Main Information</Typography>
                {/* Content */}
              </Card>
            
              <Card sx={{ p: { xs: 4, md: 6 } }}>
                <Typography variant="h4" gutterBottom>Additional Details</Typography>
                {/* Content */}
              </Card>
            </Stack>
          </Grid>
        
          <Grid item xs={12} md={4}>
            <Card sx={{ p: { xs: 4, md: 6 } }}>
              <Typography variant="h5" gutterBottom>Quick Actions</Typography>
              {/* Actions */}
            </Card>
          </Grid>
        </Grid>
      </Stack>
    </Box>
  );
}
```

### Template 3: Form View

**Use for**: Create booking, Edit customer, Add court

```tsx
function FeatureFormPage() {
  const theme = useTheme();
  const isMobile = useMediaQuery(theme.breakpoints.down('md'));
  
  return (
    <Box sx={{ p: { xs: 3, md: 6 } }}>
      <Stack spacing={4}>
        {/* Page Header */}
        <Typography variant="h2">[Create/Edit] [Item]</Typography>
      
        {/* Form Card */}
        <Card sx={{ p: { xs: 4, md: 6 } }}>
          <Stack spacing={3}>
            <Typography variant="h5">Basic Information</Typography>
          
            <TextField
              fullWidth
              label="Field Name"
              size={isMobile ? 'medium' : 'small'}
            />
          
            <Grid container spacing={2}>
              <Grid item xs={12} md={6}>
                <TextField fullWidth label="Field 1" size={isMobile ? 'medium' : 'small'} />
              </Grid>
              <Grid item xs={12} md={6}>
                <TextField fullWidth label="Field 2" size={isMobile ? 'medium' : 'small'} />
              </Grid>
            </Grid>
          
            <TextField
              fullWidth
              label="Description"
              multiline
              rows={4}
              size={isMobile ? 'medium' : 'small'}
            />
          </Stack>
        </Card>
      
        {/* Actions */}
        <Stack 
          direction={{ xs: 'column-reverse', md: 'row' }} 
          spacing={2}
          justifyContent="flex-end"
        >
          <Button variant="text" fullWidth={isMobile}>Cancel</Button>
          <Button variant="contained" fullWidth={isMobile} size={isMobile ? 'large' : 'medium'}>
            Save [Item]
          </Button>
        </Stack>
      </Stack>
    </Box>
  );
}
```

---

## 6. Content Density Guidelines

### Mobile (<900px)

- **Spacing**: More generous (easier touch targets)
- **Typography**: Larger sizes (16px minimum for inputs)
- **Buttons**: Large size (48px height)
- **Cards**: Full-width, stacked vertically
- **Layout**: Single column

### Desktop (≥900px)

- **Spacing**: Compact (efficient use of space)
- **Typography**: Standard sizes (14px for inputs)
- **Buttons**: Medium size (40px height)
- **Cards**: Multi-column grid possible
- **Layout**: Multi-column, side-by-side

### Density Levels

**Low Density (Mobile Default)**

```tsx
<Stack spacing={4}>
  <Card sx={{ p: 6 }}>
    <Stack spacing={3}>
      <TextField size="medium" />
      <Button size="large" />
    </Stack>
  </Card>
</Stack>
```

**Medium Density (Desktop Default)**

```tsx
<Stack spacing={3}>
  <Card sx={{ p: 4 }}>
    <Stack spacing={2}>
      <TextField size="small" />
      <Button size="medium" />
    </Stack>
  </Card>
</Stack>
```

**High Density (Desktop Tables)**

```tsx
<Table size="small">
  <TableRow sx={{ height: 32 }}>
    <TableCell sx={{ p: 1 }}>
      <Button size="small" />
    </TableCell>
  </TableRow>
</Table>
```

---

## 7. Quick Decision Checklist

When creating a new feature UI, ask:

### Title

- [ ] What level is this page? (Dashboard → h1, Feature → h2, Detail → h3)
- [ ] Does the title clearly describe the feature?
- [ ] Is it concise (2-4 words max)?

### Content

- [ ] Are text sizes appropriate for mobile (16px min for inputs)?
- [ ] Is body text using body1 (16px) or body2 (14px)?
- [ ] Are labels using overline variant?
- [ ] Are status indicators using caption variant?

### Cards

- [ ] Does this content need a card? (grouped info = yes, single element = no)
- [ ] Is padding responsive? (`p: { xs: 4, md: 6 }`)
- [ ] Is internal spacing consistent? (`spacing={3}`)
- [ ] Does it have a clear title (h5)?

### Actions

- [ ] Are buttons responsive? (`size={isMobile ? 'large' : 'medium'}`)
- [ ] Is button hierarchy clear? (contained > outlined > text)
- [ ] Are destructive actions using error color?
- [ ] Is layout responsive? (vertical mobile, horizontal desktop)
- [ ] Are mobile buttons full-width?

### Layout

- [ ] Is page padding responsive? (`p: { xs: 3, md: 6 }`)
- [ ] Are sections spaced properly? (`spacing={{ xs: 4, md: 6 }}`)
- [ ] Does mobile use single column?
- [ ] Does desktop use multi-column where appropriate?
- [ ] Are touch targets minimum 48×48px on mobile?

---

## Examples by Feature Type

### Booking Management

- **Page Title**: h2 - "Court Bookings"
- **Card Title**: h5 - "Court 1"
- **Content**: body1 for customer names, body2 for times
- **Actions**: "View Details" (outlined), "Cancel Booking" (error)
- **Layout**: Cards on mobile, table on desktop

### Customer Management

- **Page Title**: h2 - "Customers"
- **Card Title**: h5 - Customer name
- **Content**: body1 for contact info, caption for member ID
- **Actions**: "Edit" (outlined), "View History" (contained)
- **Layout**: Grid of cards (1 col mobile, 2-3 cols desktop)

### Court Management

- **Page Title**: h2 - "Courts"
- **Card Title**: h5 - "Court [Number]"
- **Content**: body2 for status, caption for last maintenance
- **Actions**: "Edit Court" (outlined), "Mark Maintenance" (warning)
- **Layout**: Grid of cards with status chips

---

## Related Workflows

- `/create-theme-compliant-ui` - Detailed component examples
- `/create-route-with-search-params` - Route setup with filters
- `/create-feature-module` - Complete feature structure

## Related Skills

- `ui-theme` - Theme tokens and rules
- `feature` - Feature-first architecture
- `form` - Form handling patterns
