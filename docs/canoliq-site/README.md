# canoLiq documentation site

The published documentation for canoLiq, built with Docusaurus. Content lives
in `docs/` as MDX; everything else here is theme and configuration.

## Run it locally

```bash
npm ci
npm start
```

`npm run build` produces the static site in `build/`. The build **fails on broken
links** (`onBrokenLinks: 'throw'`), so run it before opening a pull request.

## Who this site is for

Users and operators of the chain. Protocol-developer material is deliberately
confined to the **Protocol Reference** section at the bottom of the sidebar.
When adding a page, ask which of the two tracks it serves; if the answer is
"whoever is reading the plugin source", it belongs either in Protocol Reference
or in `plugin/go/canoliq/README.md`, which is the operator/developer runbook.

## Layout

| Path | Holds |
|---|---|
| `docs/tutorials/users/` | Step-by-step walkthroughs for people holding CNPY |
| `docs/tutorials/operators/` | Node, committee, monitoring, and governance runbooks |
| `docs/start-here/`, `docs/concepts/` | Plain-language introductions |
| `docs/tokenomics/`, `docs/governance/` | Economic and voting rules |
| `docs/advanced/` | Operational subsystems, shown as "Operations" |
| `docs/proto/` | Raw wire formats, shown under "Protocol Reference" |
| `sidebars.ts` | Sidebar order, curated by hand |
| `src/css/custom.css` | Brand palette |
| `static/img/` | Brand assets |

## Conventions

- **Numbers come from the code.** Cross-check any parameter against
  `DefaultParams()` and `defaultGovernanceTiers()` in
  `plugin/go/canoliq/config.go`. Protocol figures trace to the v1.2 papers.
- **Amounts are base units.** uCNPY, uCCNPY, uCPLQ, all six-decimal. Say so
  whenever a number appears in a command.
- **Tutorials must be runnable.** Every command in `docs/tutorials/` should be
  executable against localnet (`make docker/up-fast`) and produce the stated
  output. Verify before committing; docs here have been wrong on localnet before.
- **Do not document what is not built.** Mark partial features as partial.

## Brand

| Token | Hex | Use |
|---|---|---|
| Periwinkle | `#9BACFC` | Dark-mode primary, wordmark, ribbon |
| Deep indigo | `#1C1947` | Dark ground, hero |
| Indigo-blue | `#475FD1` | Light-mode primary (periwinkle is only 2.16:1 on white) |

Assets in `static/img/` are derived from the source logo artwork. Replace them
with vector originals when those are available.
