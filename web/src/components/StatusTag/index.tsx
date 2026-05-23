import styles from './StatusTag.less';

type Status = 'pending' | 'deploying' | 'running' | 'error';

interface StatusTagProps {
  status: string;
}

const statusConfig: Record<Status, { label: string; className: string }> = {
  pending: { label: 'pending', className: styles.pending },
  deploying: { label: 'deploying', className: styles.deploying },
  running: { label: 'running', className: styles.running },
  error: { label: 'error', className: styles.error },
};

export default function StatusTag({ status }: StatusTagProps) {
  const config = statusConfig[status as Status] || {
    label: status,
    className: styles.default,
  };

  return <span className={`${styles.tag} ${config.className}`}>{config.label}</span>;
}