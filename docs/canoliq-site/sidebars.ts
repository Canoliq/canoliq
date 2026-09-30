import type {SidebarsConfig} from '@docusaurus/plugin-content-docs';

/**
 * Ordered by audience, not by subsystem. Users and operators come first and
 * meet tutorials before reference; protocol internals are deliberately walled
 * off at the bottom so the default reading path never lands in raw protobuf.
 */
const sidebars: SidebarsConfig = {
  canoliqSidebar: [
    'intro',
    'network-status',
    {
      type: 'category',
      label: 'Start Here',
      link: {type: 'doc', id: 'start-here/what-is-canopy'},
      items: [
        'start-here/what-is-canopy',
        'start-here/what-is-liquid-staking',
        'start-here/reading-guide',
      ],
    },
    {
      type: 'category',
      label: 'Core Concepts',
      link: {type: 'doc', id: 'concepts/how-it-works'},
      items: [
        'concepts/how-it-works',
        'concepts/two-tokens',
        'concepts/glossary',
      ],
    },
    {
      type: 'category',
      label: 'Tutorials: Users',
      link: {type: 'doc', id: 'tutorials/overview'},
      items: [
        'tutorials/users/setup',
        'tutorials/users/deposit',
        'tutorials/users/track-your-position',
        'tutorials/users/transfer',
        'tutorials/users/redeem',
        'tutorials/users/otc-lock',
        'tutorials/users/cplq-and-voting',
      ],
    },
    {
      type: 'category',
      label: 'Tutorials: Operators',
      link: {type: 'doc', id: 'tutorials/operators/run-a-node'},
      items: [
        'tutorials/operators/run-a-node',
        'tutorials/operators/join-the-committee',
        'tutorials/operators/stake-ownership',
        'tutorials/operators/promote-to-testnet',
        'tutorials/operators/monitoring',
        'tutorials/operators/alerts',
        'tutorials/operators/governance-proposal',
      ],
    },
    {
      type: 'category',
      label: 'Tokenomics',
      link: {type: 'doc', id: 'tokenomics/overview'},
      items: [
        'tokenomics/overview',
        'tokenomics/fee-structure',
        'tokenomics/vote-escrow',
        'tokenomics/otc-lock-program',
        'tokenomics/vesting',
      ],
    },
    {
      type: 'category',
      label: 'Governance',
      link: {type: 'doc', id: 'governance/overview'},
      items: [
        'governance/overview',
        'governance/governance-tiers',
        'governance/proposals',
        'governance/tally-execution',
      ],
    },
    {
      type: 'category',
      label: 'Operations',
      link: {type: 'doc', id: 'advanced/tvl-cap'},
      items: [
        'advanced/tvl-cap',
        'advanced/insurance',
        'advanced/restaking',
        'advanced/alerts',
        'advanced/treasury',
        'advanced/buyback',
        'advanced/otc-lock-accounting',
        'advanced/autonomy-graduation',
      ],
    },
    {
      type: 'category',
      label: 'Reference',
      link: {type: 'doc', id: 'transactions/overview'},
      items: [
        'transactions/overview',
        'transactions/reference',
        'transactions/deposit-redeem',
        'transactions/cplq-operations',
        'transactions/otc-locks',
        'getting-started/overview',
        'getting-started/building',
        'getting-started/committee-setup',
        'api/overview',
        'api/endpoints',
      ],
    },
    {
      type: 'category',
      label: 'Protocol Reference',
      link: {type: 'generated-index', description:
        'Wire formats and state layout. You do not need any of this to use or operate canoLiq — it is here for people reading or extending the plugin source.'},
      items: [
        'proto/messages',
        'proto/types',
        'advanced/state-keys',
      ],
    },
  ],
};

export default sidebars;
