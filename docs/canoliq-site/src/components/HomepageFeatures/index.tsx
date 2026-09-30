import type {ReactNode} from 'react';
import clsx from 'clsx';
import Link from '@docusaurus/Link';
import Heading from '@theme/Heading';
import styles from './styles.module.css';

type FeatureItem = {
  title: string;
  icon: ReactNode;
  description: ReactNode;
  to: string;
  cta: string;
};

/* Simple geometric marks drawn in the brand's rounded-stroke idiom. */
const Liquid = (
  <svg viewBox="0 0 48 48" role="img" aria-hidden="true">
    <circle cx="24" cy="24" r="17" />
    <path d="M9 26c5 0 5 4 10 4s5-4 10-4 5 4 10 4" />
  </svg>
);

const Compound = (
  <svg viewBox="0 0 48 48" role="img" aria-hidden="true">
    <path d="M8 36c8 0 12-6 16-12s8-12 16-12" />
    <path d="M32 12h8v8" />
  </svg>
);

const Govern = (
  <svg viewBox="0 0 48 48" role="img" aria-hidden="true">
    <path d="M24 6l14 7v10c0 10-6 16-14 19-8-3-14-9-14-19V13z" />
    <path d="M18 24l4 4 9-9" />
  </svg>
);

const FeatureList: FeatureItem[] = [
  {
    title: 'Staked, but still liquid',
    icon: Liquid,
    description: (
      <>
        Deposit CNPY and receive <strong>cCNPY</strong>. Your stake keeps
        working on a Canopy committee while the receipt token stays in your
        wallet, transferable at any time.
      </>
    ),
    to: '/docs/concepts/how-it-works',
    cta: 'How it works',
  },
  {
    title: 'Rewards compound on their own',
    icon: Compound,
    description: (
      <>
        There is nothing to claim. Rewards raise the cCNPY exchange rate, so
        each token is worth more CNPY over time. Holders keep{' '}
        <strong>92.8%</strong> of rewards received.
      </>
    ),
    to: '/docs/tokenomics/fee-structure',
    cta: 'See the fee split',
  },
  {
    title: 'Governed by CPLQ holders',
    icon: Govern,
    description: (
      <>
        Fees, treasury spending, and upgrades all run through proposals. Lock
        CPLQ for up to 24 months to multiply your voting weight up to{' '}
        <strong>4&times;</strong>.
      </>
    ),
    to: '/docs/governance/overview',
    cta: 'Read about governance',
  },
];

function Feature({title, icon, description, to, cta}: FeatureItem) {
  return (
    <div className={clsx('col col--4')}>
      <div className={styles.card}>
        <div className={styles.featureIcon}>{icon}</div>
        <Heading as="h3">{title}</Heading>
        <p>{description}</p>
        <Link to={to}>{cta} &rarr;</Link>
      </div>
    </div>
  );
}

export default function HomepageFeatures(): ReactNode {
  return (
    <section className={styles.features}>
      <div className="container">
        <div className="row">
          {FeatureList.map((props, idx) => (
            <Feature key={idx} {...props} />
          ))}
        </div>
      </div>
    </section>
  );
}
