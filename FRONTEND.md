# FRONTEND.md — Brainiac Design System

Style guide for building UI that matches the Brainiac app. Copy tokens and component recipes verbatim; deviate only with a documented reason.

---

## 1. Design principles

1. **Graph-first, friendly-serious.** The product is about learning paths. Visual language is playful (pills, thick strokes, offset shadows) but the information is dense and legible.
2. **Purple is structure, orange is action.** `--p` paints surfaces, icons, and state; `--a` is reserved for the primary call-to-action, focus and selection rings, and "special" markers (passes, attestations).
3. **Flat by default, one hard shadow.** No soft blurry shadows. The single exception is the "hard offset" shadow (`8px 8px 0 var(--a)`) used for the pass card.
4. **Pills everywhere.** Buttons, tags, filters, coins, nav links, badges — all fully rounded (`border-radius: 99px`). Containers are `20px`.
5. **Contrast over decoration.** Body copy is either `--ink` or `--mut`. Never invent a third grey.
6. **Content never shifts on route change.** Views are swapped in place; scroll resets to top.
7. **Categories get their own colours, and only categories.** The `--t-*` track colours tell data categories apart (for example, learning tracks on the graph). They never style buttons, nav, or other chrome.
8. **Direct manipulation stays calm.** Graphs can be dragged, panned and zoomed. Motion is limited to physics settling and short hover/selection transitions (`.14-.15s`). No decorative animation.

---

## 2. Design tokens

Drop this into every page. These values are canonical — the tokens below override anything older in legacy files.

```css
:root {
  /* Brand */
  --p: #5941a9;      /* primary — structure, surfaces, icons, active state */
  --s: #888ef0;      /* secondary — completed/done, decorative fills */
  --a: #ee8434;      /* accent — primary CTA, focus ring, highlight */

  /* Neutrals */
  --bg: #ffffff;
  --surf: #f4f2fb;
  --ink: #000000;
  --mut: #57547a;
  --line: #dedaf0;
  --on-p: #fff;      /* label colour on --p */
  --on-a: #000;      /* label colour on --a; theme-invariant (--ink flips to white in dark mode) */

  /* Categorical: graph tracks. Theme-invariant, data only (see "Track colours") */
  --t-1: #ee8434;
  --t-2:  #888ef0;
  --t-3:   #5941a9;
  --t-4: #2fa98a;
  --t-5:  #d4608a;
  --t-6:  #3C93A9;
  --t-7:  #FECF0E;
  --t-8:  #D6595A;

  /* Shape */
  --r-pill: 99px;
  --r-lg: 20px;      /* cards, panels, images */
  --r-md: 14px;
  --r-sm: 12px;
  --r-xs: 8px;

  /* Elevation (hard, never blurred) */
  --shadow-pop: 8px 8px 0 var(--a);

  /* Rhythm */
  --gap: 20px;
  --pad-card: 26px;
  --pad-panel: 24px;
  --section-y: 72px;
  --maxw: 1120px;
  --gutter: 24px;
  --nav-h: 64px;

  color-scheme: light;
  box-sizing: border-box;
  padding-top: env(safe-area-inset-top, 0px);
  padding-bottom: env(safe-area-inset-bottom, 0px);
}

@media (prefers-color-scheme: dark) {
  :root:not([data-theme="light"]) {
    --bg: #110e22;
    --surf: #1b1735;
    --ink: #ffffff;
    --mut: #b8b4de;
    --line: #342e5e;
    /* --p, --s, --a stay identical in both themes by design */
    color-scheme: dark;
  }
}
```

**Rules**

- Never hard-code a hex value in a component. Always reference a token.
- `--p`, `--s`, `--a`, `--on-a` and every `--t-*` are theme-invariant. Do not re-map them in dark mode.
- Text on `--p` is always `--on-p`. Text on `--a` is always `--on-a` (never `--ink`, which turns white in dark mode). This pairing is accessible; don't swap it.
- Respect manual override: `data-theme="light"` on `<html>` must force light even under a dark OS preference.

### Logo

Preferred logo images are `./web/public/assets/img/logo/logo-varied-purple.svg` and `./web/public/assets/img/logo/favicon-varied-purple.svg` (For favicon only). Same logo images are used regardless of color scheme.

Logo is followed by brand title "Brainiac" in Bricolage Grotesque 1.3 rem. There is a 8px gap between logo and text, logo height is 46px.

Logo with text should always be wrapped into `a` tag, linking to the main page.

### Track colours

A five-track map needs five distinguishable fills, and the core palette has only three hues. The `--t-*` tokens extend it for categorical data.

| Token | Hex | Track | Relation to core palette | Glyph / label on the fill |
|---|---|---|---|---|
| `--t-1` | `#ee8434` | Graph item 1 | same value as `--a` | `--on-a` (black) |
| `--t-2` | `#888ef0` | Graph item 2 | same value as `--s` | `#fff` |
| `--t-3` | `#5941a9` | Graph item 3 | same value as `--p` | `#fff` |
| `--t-4` | `#2fa98a` | Graph item 4 | new (teal) | `#fff` |
| `--t-5` | `#d4608a` | Graph item 5 | new (rose) | `#fff` |
| `--t-6` | `#3c93a9` | Graph item 6 | new (ocean) `--p` | `#fff` |
| `--t-7` | `#fecf0e` | Graph item 7 | new (yellow) | `#--on-a` (black) |
| `--t-8` | `#d6595a` | Graph item 8 | new (red) | `#fff` |

- Use `--t-*` only to distinguish categories in data views (graph discs, track tags, legend swatches). Never for buttons, nav, or surfaces.
- Colors `--t-*` are also used for built-in avatar images.
- Reference tokens, not hex, wherever CSS allows it. Where a value must be passed from JS (for example `color: 'var(--t-node)'`), pass the `var()` string, not the hex.
- Keep the three aliased values in sync with `--a`, `--s` and `--p` if the brand palette ever changes.
- Adding a sixth track means adding a new `--t-*` token here first; don't invent one inline.

### Avatar images

Users have random avatars from `./web/public/assets/img/avatar/`, all files have colors #5941a9 (bg) a #ffffff (lines). Background colors for those avatars should be set randomly from `--t-*`.
Preferred seed of a default avatar should be the user id (which uses uuid v4).


### Header

Header is composed of Logo (left), links (main, courses, profile, items; also on the left or in the middle), on the right there are coin balance button, level bar, icons for messenger, notifications and account button. 

### Allowed literals

Components reference tokens, with these exceptions: `#fff` for the logo plate and for glyph/label strokes on coloured fills, and `#fbb77d` for the inner ring of the coin disc.

---

## 3. Typography

**Family:** `"Bricolage Grotesque"`, then `system-ui, -apple-system, "Segoe UI", sans-serif`.

```html
<link rel="preconnect" href="https://fonts.googleapis.com">
<link href="https://fonts.googleapis.com/css2?family=Bricolage+Grotesque:opsz,wght@12..96,400..800&display=swap" rel="stylesheet">
```

| Element | Size | Weight | Tracking | Leading |
|---|---|---|---|---|
| `h1` display | `clamp(2.4rem, 6vw, 4.4rem)` | 800 | `-.03em` | `1` |
| `h2` section | `clamp(1.6rem, 3.5vw, 2.4rem)` | 800 | `-.02em` | `1.1` |
| `h3` card title | `1.15rem` | 800 | normal | default |
| Body | `1rem` | 400 | normal | `1.5` |
| Lead paragraph | `1.2rem` | 400 | normal | `1.5` |
| Eyebrow (above h1/h2) | `.85rem` | 600 | normal | default |
| Side-panel heading (`h2.side-h`) | `1.4rem` | 800 | `-.01em` | `1.2` |
| Small / meta | `.85–.9rem` | 400 | normal | default |
| Stat number | `2–2.4rem` | 800 | normal | `1` |

- Headings use `margin: 0 0 12–20px`. Paragraphs use `margin: 0 0 14px`.
- Secondary copy uses `.mut` (`color: var(--mut)`), never a lighter opacity of `--ink`.
- Numerals in stats, balances, and progress are always `font-weight: 800`.
- Lead paragraphs in page headers are capped at `max-width: 56ch`.

```css
body{font-family:"Bricolage Grotesque",system-ui,-apple-system,"Segoe UI",sans-serif;
     line-height:1.5;background:var(--bg);color:var(--ink);margin:0}
h1{font-size:clamp(2.4rem,6vw,4.4rem);line-height:1;letter-spacing:-.03em;font-weight:800;margin:0 0 20px}
h2{font-size:clamp(1.6rem,3.5vw,2.4rem);line-height:1.1;letter-spacing:-.02em;font-weight:800;margin:0 0 12px}
h3{margin:0 0 6px;font-size:1.15rem;font-weight:800}
p{margin:0 0 14px}
.mut{color:var(--mut)}
```

---

## 4. Layout

```css
.wrap{max-width:var(--maxw);margin:0 auto;padding:0 var(--gutter)}
section.blk{padding:var(--section-y) 0}
.grid3{display:grid;grid-template-columns:repeat(3,1fr);gap:var(--gap);margin-top:32px}
.split{display:grid;grid-template-columns:1fr 1fr;gap:48px;align-items:center}
.cols{display:grid;grid-template-columns:1.5fr 1fr;gap:28px;padding:36px 0 64px}

/* page header: copy left, stat block right */
.head{display:flex;align-items:flex-end;gap:28px;flex-wrap:wrap;padding:48px 0 0}
.head .lead{font-size:1.2rem;line-height:1.5;max-width:56ch;margin:0}
.head h1{margin:0 0 12px}
.eyebrow{font-size:.85rem;font-weight:600;margin:0 0 10px}

/* tool layout: fluid canvas + fixed sidebar */
.graph-layout{display:grid;grid-template-columns:minmax(0,1fr) 340px;gap:24px;
              padding:28px 0 64px;align-items:start}
.side{position:sticky;top:calc(var(--nav-h) + 24px)}
```

Base resets that every page carries:

```css
*,*::before,*::after{box-sizing:border-box}
body{-webkit-font-smoothing:antialiased}
a{color:inherit;text-decoration:none}
:focus-visible{outline:3px solid var(--a);outline-offset:2px}
.sr-only{position:absolute;width:1px;height:1px;padding:0;margin:-1px;overflow:hidden;
         clip:rect(0 0 0 0);white-space:nowrap;border:0}
```

- Content width: `1120px` max, `24px` gutters.
- Vertical section rhythm: `72px` top and bottom.
- Two-column "split" sections pair copy with a visual/object; on mobile they stack.
- Sidebar layouts use `1.5fr 1fr` (main content left, actions/status right).
- Auto-fill card grids: `repeat(auto-fill, minmax(200px,1fr))`.
- Page header (`.head`): eyebrow, `h1` and lead on the left; a headline stat (`.bal`) pushed right with `margin-left:auto`.
- Tool layouts (interactive canvas plus detail): a fluid `minmax(0,1fr)` main column and a fixed `340px` sidebar. `minmax(0,…)` lets the canvas shrink. The sidebar is sticky at `top: calc(var(--nav-h) + 24px)`.

**Responsive breakpoint — a single `820px` query:**

```css
@media (max-width:820px){
  .hero .wrap,.split,.cols{grid-template-columns:1fr}
  .grid3{grid-template-columns:1fr}
  .stats{margin-left:0}
  nav .wrap{gap:12px}
  .logo svg{height:26px}
  nav .links a{padding:8px 10px}
  .bal{margin-left:0;width:100%}

  /* tool layouts */
  .head{padding:32px 0 0}
  .graph-layout{grid-template-columns:1fr;padding:20px 0 48px}
  .side{position:static}
  .canvas{height:clamp(380px,62vh,520px)}
  .stage-bar .hint{order:2;flex-basis:100%;font-size:.85rem}
  .zoom{order:1;margin-left:auto}
}
```

Mobile rules: stack all multi-column layouts, shrink logo to `26px` height, tighten nav link padding, let full-width blocks (balance card) span `100%`. On tool layouts the sidebar drops below the canvas and stops being sticky, the canvas shrinks to `clamp(380px,62vh,520px)`, the zoom controls move to the right, and the hint text wraps onto its own line below.

---

## 5. Components

### 5.1 Buttons

```css
.btn{font:inherit;font-weight:600;border:0;border-radius:var(--r-pill);padding:11px 22px;
     background:var(--a);color:var(--on-a);cursor:pointer;
     display:inline-flex;align-items:center;justify-content:center;min-height:46px}
.btn.alt{background:transparent;border:2px solid var(--p);color:var(--p)}
.btn.sm{padding:7px 16px;font-size:.9rem;min-height:44px}
.btn:disabled{opacity:.5;cursor:not-allowed}
.zbtn{width:44px;height:44px;padding:0;font-size:1.3rem;font-weight:700;line-height:1} /* icon-only */
```

- **Primary** (`.btn`): accent fill, black label. One per view/section.
- **Secondary** (`.btn.alt`): transparent, 2px primary outline, primary label.
- **Small** (`.btn.sm`): inline/card-level actions.
- **Disabled**: use the `disabled` attribute with `.btn:disabled` (`opacity: .5; cursor: not-allowed;`). Do not change the color.
- **Icon-only** (`.btn.alt.sm.zbtn`): `44×44px` circle. Glyphs are `18px` inline SVG with `stroke="currentColor"`, or a text character such as `+` or `−`.
- **Pending**: while an action runs, swap the label to a present participle with an ellipsis ("Opening…"), disable the button, then restore both.
- Button text colour on `--a` is `--on-a`.
- Never use a border-radius other than the pill for buttons.
- Icon-only buttons must carry an `aria-label`.

### 5.2 Navigation

```css
nav{position:sticky;top:env(safe-area-inset-top,0px);z-index:5;
    background:var(--bg);border-bottom:1px solid var(--line)}
nav .wrap{display:flex;align-items:center;gap:28px;height:var(--nav-h)}
.logo{display:flex;align-items:center;background:#fff;border-radius:var(--r-sm);padding:4px 8px}
.logo svg{height:30px;width:auto}
nav .links{display:flex;gap:6px;flex:1}
nav .links a{padding:8px 14px;border-radius:var(--r-pill);font-weight:600;color:var(--mut)}
nav .links a[aria-current=page]{background:var(--p);color:#fff}
```

- Logo sits on a white chip — the mark is purple and needs a light plate in both themes.
- Logo mark is a four-node path glyph in the brand colours (`--p`, `--s`, `--p`, `--a`):

```css
.logo .e{fill:none;stroke:var(--p);stroke-width:3.5;stroke-linecap:round;stroke-linejoin:round}
.logo .n-p{fill:var(--p)} .logo .n-s{fill:var(--s)} .logo .n-a{fill:var(--a)}
```

```html
<a class="logo" href="#landing" aria-label="Brainiac home">
  <svg viewBox="0 0 44 32" aria-hidden="true">
    <path class="e" d="M7 23 L17 9 L27 19 L37 7"/>
    <circle class="n-p" cx="7"  cy="23" r="4.5"/><circle class="n-s" cx="17" cy="9"  r="4.5"/>
    <circle class="n-p" cx="27" cy="19" r="4.5"/><circle class="n-a" cx="37" cy="7"  r="4.5"/>
  </svg>
</a>
```

- The active page is marked with `aria-current="page"` and a primary pill fill.
- Right cluster: guest CTA (`land-only`) or coin balance + avatar (`app-only`).

### 5.3 Cards, panels, tags

```css
.card{background:var(--surf);border-radius:var(--r-lg);padding:var(--pad-card)}
.card .ic{width:44px;height:44px;border-radius:var(--r-sm);background:var(--p);color:#fff;
          display:grid;place-items:center;margin-bottom:16px}
.panel{border:1px solid var(--line);border-radius:var(--r-lg);padding:var(--pad-panel);
       margin-bottom:24px;background:var(--bg)}
.panel.emphasis{border:2px solid var(--a)}          /* assessment pass / primary action panel */
.pass{background:var(--bg);border:2px solid var(--line);border-radius:var(--r-lg);
      padding:22px;box-shadow:var(--shadow-pop)}
.tag{display:inline-block;background:var(--a);color:var(--on-a);border-radius:var(--r-pill);
     padding:2px 12px;font-size:.8rem;font-weight:800;margin-bottom:12px}
```

- **Card** = tinted `--surf`, borderless, icon block on top. Use for feature/benefit triads.
- **Panel** = bordered, on `--bg`. Use for stacked content in dense layouts.
- **Emphasis panel** = 2px accent border, no fill. Only for the one thing the user should act on now.
- **Pass card** = the only element allowed a hard drop shadow.
- **Tag** = accent pill, `.8rem`, weight 800. Secondary tags may use `--s` fill with white text.
- **Track tag** = same pill with `background: var(--t-*)` and the glyph/label pairing from the track-colour table (white, except HTML which uses `--on-a`).
- **Sidebar stack** = several `.panel`s in a column, each opening with an `h3` (for example "States", "Tracks"). The selection-detail panel at the top is a `.panel` with `aria-live="polite"`.

### 5.4 Pills: coin balance & identity

```css
.coin{display:inline-flex;align-items:center;gap:8px;background:var(--surf);
      border-radius:var(--r-pill);padding:6px 14px 6px 8px;font-weight:800}
.coin i{width:22px;height:22px;border-radius:50%;background:var(--a);
        box-shadow:inset 0 0 0 4px #fbb77d}
.me{width:36px;height:36px;border-radius:50%;background:var(--s);color:#fff;
    display:grid;place-items:center;font-weight:800}
```

- Coin balance is a pill with a solid accent disc inside. Same component on landing, profile, inventory.
- Avatar is a `--s` circle with 2-letter initials, weight 800. Give it `title="Signed in as <name>"`.
- In the nav the coin number is followed by `<span class="sr-only">coins</span>` so screen readers get the unit.
- Large profile avatar: `112px`, `2.6rem` initials, double ring via `box-shadow: 0 0 0 6px var(--bg), 0 0 0 10px var(--a)`.

### 5.5 Progress & status

```css
.bar{height:10px;border-radius:9px;background:var(--line);overflow:hidden;margin:10px 0 6px}
.bar span{display:block;height:100%;background:var(--p)}      /* use var(--s) for secondary tracks */

.subj{display:flex;align-items:center;gap:12px;padding:11px 0;border-top:1px solid var(--line)}
.dot{width:14px;height:14px;border-radius:50%;flex:none;border:3px solid var(--p)}
.dot.done{background:var(--p)}
.dot.cur{border-color:var(--a)}
.dot.lock{border-color:var(--line);border-style:dashed}
.subj span.g{margin-left:auto;color:var(--mut);font-size:.9rem}
```

Status vocabulary — use consistently for subjects, steps, and graph nodes:

| State | Visual | Label pattern |
|---|---|---|
| Done | filled `--p` dot; on the graph, a `--s` check badge | "Completed", "Assessment, 88%" |
| Current | `--a` ring | "In progress" |
| Open | plain `--p` ring | "Available" |
| Locked | dashed `--line` dot; on the graph, a dimmed dashed node with a padlock badge | "Locked until X and Y" |

- Never signal state by colour alone. Pair it with a shape cue (check badge, padlock, dashed outline).
- Legend swatches mirror the graph (Done = `--s` fill); status dots in lists use the dot rules above (Done = `--p` fill).
- On the interactive graph the `--a` ring means *selected* or *keyboard-focused*, not "Current".
Right-aligned meta text is always `--mut` at `.9rem`.

### 5.6 Tabs / filters

```css
.tabs{display:flex;gap:8px;flex-wrap:wrap;margin-bottom:24px}
.tabs button{border:2px solid var(--line);background:var(--bg);color:var(--ink);
             border-radius:var(--r-pill);padding:8px 18px;font-weight:600}
.tabs button[aria-pressed=true]{background:var(--p);border-color:var(--p);color:#fff}
```

Toggle state is expressed with `aria-pressed`, not a class. Wrap in `role="group"` with an `aria-label`.

### 5.7 Item grid (inventory)

```css
.items{display:grid;grid-template-columns:repeat(auto-fill,minmax(200px,1fr));gap:18px;padding-bottom:64px}
.item{border:1px solid var(--line);border-radius:var(--r-lg);padding:16px;background:var(--bg);
      display:flex;flex-direction:column;gap:6px}
.art{height:120px;border-radius:var(--r-md);display:grid;place-items:center;margin-bottom:8px}
.art svg{width:64px;height:64px}
.item .btn{margin-top:auto}
.empty{grid-column:1/-1;color:var(--mut);padding:40px 0}
```

- Art tile is a solid brand-color block (`--p`, `--s`, or `--a`) with a 64px white/orange line glyph inside.
- Action button is pinned to the bottom with `margin-top:auto`.
- Empty state spans the full grid and uses `--mut`.

### 5.8 Balance / hero stat block

```css
.bal{margin-left:auto;background:var(--p);color:#fff;border-radius:var(--r-lg);
     padding:18px 26px;min-width:220px}
.bal b{font-size:2.4rem;display:block;line-height:1;font-weight:800}
.bal>span{font-size:.9rem;font-weight:600;display:block;margin-bottom:6px}   /* label above the number */
.bal-of{font-size:1.1rem;font-weight:600}                                    /* "/ 20" suffix */
```

```html
<div class="bal"><span>Completed</span><b><span class="bal-n">0</span><span class="bal-of"> / 20</span></b></div>
```

Headline stat cards use primary fill with white text and the number at `2.4rem` / 800. The label sits above at `.9rem` / 600. The block works for any headline count (coins, completed subjects), not only currency; a "of total" suffix uses `.bal-of`. Use `.bal>span` for the label so the nested number spans are not restyled.

### 5.9 Attendance / QR tile

```css
.qr{width:96px;height:96px;flex:none;background:conic-gradient(var(--ink) 0 25%,transparent 0 50%,
    var(--ink) 0 75%,transparent 0) 0 0/24px 24px,var(--bg);
    border:6px solid var(--bg);outline:3px solid var(--ink);border-radius:6px}
.week{display:flex;gap:6px;margin-top:14px}
.week i{flex:1;height:30px;border-radius:8px;background:var(--line)}
.week i.y{background:var(--s)}
```

Week strip: `8px` radius bars, `--line` for absent, `--s` for present. Always caption it.

### 5.10 Interactive subjects graph (SVG)

The signature visual of the app. Circular, track-coloured nodes on the page background, laid out by a force simulation the person can drag, pan and zoom. Behaviour lives in section 6.5.

```css
/* canvas sits on the site background: no tinted plate, no soft shadow */
.canvas{position:relative;height:clamp(440px,calc(100vh - 300px),740px);
        background:var(--bg);border:1px solid var(--line);border-radius:var(--r-lg);
        overflow:hidden;touch-action:none;user-select:none;-webkit-user-select:none}
#graph{display:block;width:100%;height:100%;cursor:grab}
#graph.panning{cursor:grabbing}

.stage{min-width:0}
.stage-bar{display:flex;align-items:center;gap:16px;margin-bottom:14px;flex-wrap:wrap}
.stage-bar .hint{margin:0;font-size:.9rem;flex:1 1 220px}
.zoom{display:flex;gap:8px}

/* edges: straight lines */
.edge{stroke:var(--line);stroke-width:2;fill:none}
.edge.lock{stroke-dasharray:5 6;opacity:.6}
.edge.hl{stroke:var(--p);stroke-width:3;opacity:1}
#graph.has-sel .edge:not(.hl){opacity:.3}

/* nodes */
.node{cursor:grab;outline:none}
.node.dragging{cursor:grabbing}
.node .inner{transform-box:fill-box;transform-origin:center;transition:transform .14s ease}
.node:hover .inner{transform:scale(1.07)}
.node .disc{stroke:var(--bg);stroke-width:3}
.node .ring{fill:none;stroke:var(--a);stroke-width:3.5;opacity:0;transition:opacity .15s ease}
.node.sel .ring,.node:focus-visible .ring{opacity:1}
.node .glyph{fill:none;stroke:#fff;stroke-width:2;stroke-linecap:round;stroke-linejoin:round;pointer-events:none}
.node .label{font:700 12px "Bricolage Grotesque",system-ui,sans-serif;fill:var(--mut);
             text-anchor:middle;pointer-events:none}
.node .badge{pointer-events:none}

/* track fills */
.node[data-topic="html"] .disc{fill:var(--t-html)}
.node[data-topic="css"]  .disc{fill:var(--t-css)}
.node[data-topic="js"]   .disc{fill:var(--t-js)}
.node[data-topic="node"] .disc{fill:var(--t-node)}
.node[data-topic="sql"]  .disc{fill:var(--t-sql)}
.node[data-topic="html"] .glyph{stroke:var(--on-a)}

/* state: locked */
.node[data-state="locked"] .inner{opacity:.42}
.node[data-state="locked"] .disc{stroke:var(--mut);stroke-width:2.5;stroke-dasharray:5 6}
.node[data-state="locked"]:hover .inner{transform:scale(1.03)}

/* state badges (top-right of the disc) */
.badge circle{stroke:var(--bg);stroke-width:3}
.badge-done circle{fill:var(--s)}
.badge-done path{fill:none;stroke:#fff;stroke-width:2.2;stroke-linecap:round;stroke-linejoin:round}
.badge-lock circle{fill:var(--bg);stroke:var(--line);stroke-width:2.5}
.badge-lock rect{fill:none;stroke:var(--mut);stroke-width:1.6}
.badge-lock path{fill:none;stroke:var(--mut);stroke-width:1.6;stroke-linecap:round}

/* selection dimming */
#graph.has-sel .node:not(.sel):not(.nb){opacity:.42}
```

Node anatomy (one `<g class="node">` per subject, containing `<g class="inner">`):

| Part | Spec |
|---|---|
| Disc | radius `28`, filled by `data-topic`, `3px` stroke in `--bg` so overlaps read cleanly |
| Glyph | 24×24 line icon, `translate(-12,-12)`, `2px` white stroke (HTML uses `--on-a`) |
| Label | short name, `12px` / 700, `--mut`, centred at `y = R + 20` below the disc |
| Selection ring | radius `R + 6`, `3.5px` `--a` stroke, hidden by default, shown on `.sel` and `:focus-visible` |
| Done badge | `r = 11` circle in `--s` with a white check, at `(0.72R, −0.72R)` |
| Locked badge | `r = 11` circle in `--bg` with a `--line` outline and a padlock glyph in `--mut`, same position |
| Open | no badge |

- Node data attributes: `data-topic="html|css|js|node|sql"` and `data-state="done|open|locked"`.
- Track glyphs: HTML `‹ ›` brackets, CSS slashed hash, JavaScript curly braces, Node.js hexagon, SQLite database cylinder. Add a new track's glyph with its token.
- Edges are straight lines. An edge is `.lock` when either endpoint is locked. When a node is selected, its edges get `.hl` (`--p`, `3px`) and everything not selected or adjacent drops to `.42` / `.3` opacity.
- Selection, drag and neighbour state are class toggles (`.sel`, `.nb`, `.dragging`, `.has-sel`); status for assistive tech is carried in `aria-label` (see section 7).
- Do not tint the canvas or add shadows. The graph uses the surface the rest of the page uses.

### 5.11 Detail panel (selection sidebar)

```css
.side-h{font-size:1.4rem;line-height:1.2;letter-spacing:-.01em;margin:0 0 10px}
.state-line{display:flex;align-items:center;gap:10px;margin:0 0 16px;font-size:.95rem;font-weight:600}
.state-line.is-locked{color:var(--mut);font-weight:400}
.facts{list-style:none;margin:0 0 20px;padding:0}
.facts li{display:flex;gap:12px;padding:11px 0;border-top:1px solid var(--line);font-size:.95rem}
.facts li:first-child{border-top:0;padding-top:0}
.facts b{margin-left:auto;font-weight:800;text-align:right}
.side-actions{display:grid;gap:10px}
.side-actions .btn{width:100%}
.btnrow{display:flex;gap:10px;flex-wrap:wrap}
.btnrow .btn{flex:1 1 0;min-width:118px}
```

Order inside the panel, top to bottom: track tag, `h2.side-h` title, `.mut` description, `.state-line` (status dot plus label), `.facts`, actions.

- Facts are label-left, value-right (`b`), for example "Connections 4", "Task Submitted", "Assessment 88%". Show a fact row only when the data exists.
- Primary action by state: Done, "Review subject"; Open, "Start subject"; Locked, a disabled "Locked". Secondary actions (`Task`, `Assessment`) are `.btn.alt.sm` in a `.btnrow`, disabled when locked.
- Locked state text names the prerequisites: "Locked until Forms and inputs and Flexbox layout". Join with commas and a final "and".
- Empty selection shows an eyebrow ("Nothing selected"), an `h2.side-h` ("Pick a subject") and a one-paragraph explanation.

### 5.12 Legends

```css
.legend{list-style:none;margin:0;padding:0}
.legend li{display:flex;align-items:center;gap:10px;padding:9px 0;border-top:1px solid var(--line);font-size:.95rem}
.legend li:first-child{border-top:0;padding-top:0}
.legend .sw{width:14px;height:14px;border-radius:50%;flex:none}
.legend b{margin-left:auto;font-weight:800}
.sw-done{background:var(--s)}
.sw-open{border:3px solid var(--p)}
.sw-lock{border:2px dashed var(--mut);opacity:.7}
```

- One legend per dimension ("States", "Tracks"), each in its own `.panel`. Track swatches use `background: var(--t-*)`.
- The count at the right is computed from the data, not hard-coded.

### 5.13 Hero graph (static, landing only)

A decorative, non-interactive variant used in the landing hero. It lives on `--p`; the interactive graph (5.10) lives on `--bg` and uses circular nodes. Do not mix the two styles in one view. Rules:

```css
.graph .e{fill:none;stroke:#fff;stroke-width:2.5;opacity:.55}
.graph .e.lock{stroke-dasharray:5 6;opacity:.35}
.graph .n rect{fill:#fff}
.graph .n text{font:700 13px "Bricolage Grotesque",system-ui,sans-serif;fill:var(--p)}
.graph .n.done rect{fill:var(--s)}   .graph .n.done text{fill:#fff}
.graph .n.cur rect{stroke:var(--a);stroke-width:4}
.graph .n.lock rect{fill:rgba(255,255,255,.14);stroke:rgba(255,255,255,.5);stroke-dasharray:5 5}
.graph .n.lock text{fill:#e6e1fa}
.graph .n.att rect{fill:var(--a)}    .graph .n.att text{fill:#000}
```

- Node rects: `88×44` (labels) and `110×44` (assessment), `rx: 14`.
- Edges are cubic beziers; locked edges are dashed and dimmed.
- This variant always lives on a `--p` background — that's what makes the white/orange/`--s` nodes read.
- Node classes: `done`, `cur`, `lock`, `att`. Every graph needs `role="img"` and a descriptive `aria-label` that names the nodes and their states.

---

## 6. Patterns

### 6.1 View routing (SPA, no framework)

```css
.view{display:none}
.view.on{display:block}
body[data-v=landing] .app-only,
body:not([data-v=landing]) .land-only{display:none}
```

- `body[data-v="landing|subjects|profile|inventory"]` drives everything.
- The landing page is the only public view; authenticated-only chrome (`.app-only`) and the guest CTA (`.land-only`) swap automatically.
- Hash routes: `#landing`, `#subjects`, `#profile`, `#inventory`. Unknown hashes fall back to `landing`.
- On route change: toggle `.on`, sync `aria-current`, `window.scrollTo(0,0)`.
- Every nav link must resolve to a real view. Do not ship a link whose view isn't in the page.
- A view with measured layout (the graph canvas) can't measure while `display:none`. Re-run its fit when the view becomes active.

```js
const VIEWS = ['landing', 'subjects', 'profile', 'inventory'];
function route() {
  const h = location.hash.slice(1);
  const v = VIEWS.includes(h) ? h : 'landing';
  document.body.dataset.v = v;
  document.querySelectorAll('.view').forEach(el => el.classList.toggle('on', el.id === 'v-' + v));
  document.querySelectorAll('nav .links a').forEach(a => {
    if (a.getAttribute('href') === '#' + v) a.setAttribute('aria-current', 'page');
    else a.removeAttribute('aria-current');
  });
  window.scrollTo(0, 0);
}
addEventListener('hashchange', route);
route();
```

### 6.2 Landing page rhythm

1. **Hero** — primary `--p` background, white text, 2-column: copy left, graph right. Buttons: primary + white-outline secondary (`.btn.alt` with white border/text over `--p`).
2. **Feature triad** — `h2` + muted lead + 3 `.card`s with `.ic` icons.
3. **Alternating split sections** — on `--surf` and on `--bg`; one side copy, other side a demo object (pass card, coin pill cluster).
4. **CTA band** — `.cta` on `--surf`, centered, single primary button.
5. **Footer** — `1px` top border, muted, `.9rem`.

### 6.3 Coin economy UI

- Balance always shows with thousands separators (`1,240`).
- Converting an item: add value to balance → remove item → re-render grid → update **every** `.bal-n` on the page at once.
- Non-convertible items show a disabled "Keep" button, never a hidden action.

### 6.4 Safe areas

Mobile pages must use `env(safe-area-inset-top/bottom)` on `:root` padding and on the sticky nav's `top`. Add `viewport-fit=cover` to the meta viewport.

### 6.5 Interactive graph behaviour

Applies to the graph in 5.10. All input uses pointer events, so mouse, touch and pen behave the same.

| Gesture | Behaviour |
|---|---|
| Drag a node | Pointer-captured; the node follows the pointer with the grab offset kept. Collisions push neighbours away. |
| Tap a node | A press with under `7px` of total movement counts as a tap: velocity is zeroed and the node is selected. |
| Release after a drag | The node keeps `45%` of its last velocity (a light throw), then physics settles it. |
| Drag empty canvas | Pans the view. |
| Wheel | Zooms about the cursor by `exp(-deltaY * .0015)`; the point under the cursor stays fixed. |
| `+` / `−` buttons | Zoom about the canvas centre by `×1.25` / `×0.8`. |
| Fit button | Clears the manual view and re-fits the whole graph. |
| Enter / Space on a focused node | Selects it. |
| Resize | Re-fits after `150ms`, but only if the person hasn't panned or zoomed manually. |

Zoom is clamped to `0.35–2.6`. Fit adds `44px` padding and clamps its scale to `0.35–1.25`. Its bounds include extra clearance for the badge (`14px` above) and the label (`20px` below).

Simulation constants (change them together, since they are tuned as a set):

| Constant | Value | Role |
|---|---|---|
| `R` | `28` | Disc radius |
| `REP` | `2600` | Node-node repulsion, applied as `REP / d²` |
| `LEN` | `152` | Rest length of an edge spring |
| `SPRING` | `.014` | Spring stiffness |
| `GRAV` | `.0009` | Pull toward the origin |
| `DAMP` | `.82` | Velocity damping per step |
| `MAXV` | `16` | Velocity cap |
| `MIN_D` | `2R + 16` | Hard minimum spacing, enforced by a collision pass so labels stay readable |

- **Seeding**: place each track's nodes in a small cluster (radius `72`) at an angle around an ellipse (`250 × 210`), one slot per track, starting at the top.
- **Boot**: run `320` physics steps before the first paint so the graph appears already settled, then `fit()`.
- **Loop**: fixed `16.67ms` timestep driven by `requestAnimationFrame`, at most `3` steps per frame, and `dt` clamped to `60ms` so a backgrounded tab doesn't explode on return.
- **Links**: edges are a within-track ring plus cross-track "bridges". Keep the link list as data (`[idA, idB]` pairs) beside the subject list.
- **Selection**: selecting a node highlights its edges, marks its neighbours, dims the rest, and re-renders the detail panel (5.11).

---

## 7. Accessibility (non-negotiable)

- `:focus-visible { outline: 3px solid var(--a); outline-offset: 2px; }` — never remove focus outlines.
- Interactive state uses ARIA attributes: `aria-current="page"` for nav, `aria-pressed` for filters/toggles.
- Decorative SVGs get `aria-hidden="true"`; meaningful static SVGs get `role="img"` + `aria-label`.
- Interactive SVGs are different: the `<svg>` gets `role="group"` with an `aria-label` that says what's inside and how to move ("Use Tab to move between subjects"). Each node is `role="button"` with `tabindex="0"` and an `aria-label` of the form "Title, Track track, State. Select for details."
- A node's keyboard focus shows the same `--a` ring as selection (`.node:focus-visible .ring`), and Enter and Space select it.
- Panels whose content changes on selection carry `aria-live="polite"`.
- Use `.sr-only` for text a sighted user gets from an icon or position (for example the "coins" unit, or a hidden `h2` naming the graph region).
- State is never colour-only: badges, dashes and labels carry it too.
- Disabled buttons keep `disabled` (not `pointer-events:none`) so they're skipped by tab order.
- Body text contrast: `--ink` on `--bg` and `--mut` on `--bg` both pass; do not lighten `--mut`.
- Minimum tap target `44×44px` — small buttons must be padded to reach it (`.btn.sm` has `min-height:44px`; icon-only buttons are `44×44`).
- Every view has exactly one `h1`.

---

## 8. Content & voice

- Sentence case for headings and buttons. No title case, no ALL CAPS.
- Buttons are verbs: "Start learning", "Start attempt", "Scan code", "Convert to 120 coins".
- Empty states explain the next action: "Nothing here yet. Earn more coins to fill this shelf."
- Meta labels are terse and factual: "7 of 12 subjects finished", "Assessment, 88%", "Locked until Algebra and Geometry".
- Second person, present tense. No exclamation marks.
- Eyebrows are short category labels above the `h1` ("Web development track"), set in `.mut`.
- Graph hints are terse and separated with a middle dot: "Drag nodes · scroll to zoom · drag the canvas to pan".
- Pending labels use a present participle and the ellipsis character: "Opening…".

---

## 9. Do / Don't

| Do | Don't |
|---|---|
| Use `--p` for structural fills and `--a` for the single primary action | Mix accent and primary in the same button |
| Pill every button, tag, filter, badge | Square or 8px-radius buttons |
| One hard `8px 8px 0` shadow, on the pass card only | Blurred/soft box-shadows anywhere |
| `--surf` for feature cards, `--line` border for panels | Two borders and a fill on the same element |
| 20px container radius, 99px interactive radius | Inconsistent radii across cards |
| `--mut` for all secondary text | `opacity: .7` on `--ink` |
| Stack to one column under `820px` | Introduce intermediate breakpoints |
| Use `--t-*` only to tell categories apart in data views | Use track colours on buttons, nav or surfaces |
| Pair every state colour with a shape cue (badge, dash, label) | Signal state by colour alone |
| Put the interactive graph on `--bg` with circular nodes | Put it on a tinted plate, or mix with the static `--p` hero graph |
| Pass `var(--t-*)` strings from JS | Duplicate track hex values in scripts |

---

## 10. Boilerplate

```html
<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1, viewport-fit=cover">
<title>Brainiac</title>
<link rel="preconnect" href="https://fonts.googleapis.com">
<link href="https://fonts.googleapis.com/css2?family=Bricolage+Grotesque:opsz,wght@12..96,400..800&display=swap" rel="stylesheet">
<style>
/* 1. tokens (section 2, incl. --on-a and --t-*)  2. base + resets (sections 3-4)
   3. layout (section 4)  4. components (section 5)  5. patterns (section 6)
   6. @media (max-width:820px) */
</style>
</head>
<body data-v="landing">
  <nav><!-- logo | links (Learn, Subjects, Inventory) | coin + avatar --></nav>
  <main>
    <div class="view on" id="v-landing">…</div>
    <div class="view" id="v-subjects">…</div>
    <div class="view" id="v-profile">…</div>
    <div class="view" id="v-inventory">…</div>
  </main>
  <footer><div class="wrap">© Brainiac 2026</div></footer>
</body>
</html>
```

**Checklist before shipping a page:** tokens only (no raw hex beyond the allowed literals), one `h1`, one primary button per section, focus rings intact, `820px` layout verified, dark mode verified, all icon-only buttons labeled, all SVG states labeled, every nav link resolves to a view, track colours used only for categories, graph nodes reachable and operable by keyboard, state never conveyed by colour alone.
