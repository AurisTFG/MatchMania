import React from 'react';
import styles from './Content.module.css';

export default function Content({ children }: { children: React.ReactNode }) {
  return <div className={styles.contentContainer}>{children}</div>;
}
