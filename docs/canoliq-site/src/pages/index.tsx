import type {ReactNode} from 'react';
import clsx from 'clsx';
import Link from '@docusaurus/Link';
import useBaseUrl from '@docusaurus/useBaseUrl';
import Layout from '@theme/Layout';
import HomepageFeatures from '@site/src/components/HomepageFeatures';

import styles from './index.module.css';

function HomepageHeader() {
  return (
    <header className={styles.heroBanner}>
      <img
        className={styles.ribbon}
        src={useBaseUrl('/img/logo.png')}
        alt=""
        aria-hidden="true"
      />
      <div className={clsx('container', styles.heroInner)}>
        <img
          className={styles.wordmark}
          src={useBaseUrl('/img/wordmark.png')}
          alt="canoLiq"
        />
        <p className={styles.heroSubtitle}>
          Deposit CNPY and receive cCNPY, a receipt token that earns staking
          rewards while staying transferable. Governed by CPLQ holders.
        </p>
        <div className={styles.buttons}>
          <Link
            className="button button--primary button--lg"
            to="/docs/tutorials/users/deposit">
            Make your first deposit
          </Link>
          <Link
            className="button button--secondary button--outline button--lg"
            to="/docs/start-here/what-is-liquid-staking">
            What is liquid staking?
          </Link>
        </div>
      </div>
    </header>
  );
}

export default function Home(): ReactNode {
  return (
    <Layout
      title="canoLiq Documentation"
      description="Guides and tutorials for using and operating canoLiq, the liquid staking protocol on Canopy Network.">
      <HomepageHeader />
      <main>
        <HomepageFeatures />
      </main>
    </Layout>
  );
}
