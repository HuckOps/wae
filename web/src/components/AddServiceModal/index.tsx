import { getClusters } from '@/services/api';
import { FormattedMessage, useIntl } from '@@/plugin-locale';
import { useEffect, useState } from 'react';
import styles from './AddServiceModal.less';

interface ServiceFormData {
  name: string;
  repo: string;
  domain: string;
  cluster: string;
  description: string;
}

interface FormErrors {
  name?: string;
  repo?: string;
  domain?: string;
  cluster?: string;
}

interface AddServiceModalProps {
  isOpen: boolean;
  onClose: () => void;
  onSubmit: (data: ServiceFormData) => void;
}

export default function AddServiceModal({
  isOpen,
  onClose,
  onSubmit,
}: AddServiceModalProps) {
  const intl = useIntl();

  const [formData, setFormData] = useState<ServiceFormData>({
    name: '',
    repo: '',
    domain: '',
    cluster: '',
    description: '',
  });

  const [errors, setErrors] = useState<FormErrors>({});
  const [clusters, setClusters] = useState<API.Cluster[]>([]);

  useEffect(() => {
    if (isOpen) {
      fetchClusters();
    }
  }, [isOpen]);

  const fetchClusters = async () => {
    try {
      const resp = await getClusters();
      if (resp) {
        setClusters(resp.data);
      }
    } catch (error) {
      console.error('Failed to fetch clusters');
    }
  };

  const handleChange = (key: keyof ServiceFormData, value: string) => {
    setFormData((prev) => ({ ...prev, [key]: value }));
    if (errors[key as keyof FormErrors]) {
      setErrors((prev) => ({ ...prev, [key]: undefined }));
    }
  };

  const validate = (): boolean => {
    const newErrors: FormErrors = {};

    if (!formData.name.trim()) {
      newErrors.name = intl.formatMessage({ id: 'service.nameRequired' });
    }
    if (!formData.repo.trim()) {
      newErrors.repo = intl.formatMessage({ id: 'service.repoRequired' });
    } else if (!isValidUrl(formData.repo)) {
      newErrors.repo = intl.formatMessage({ id: 'service.repoInvalid' });
    }
    if (!formData.domain.trim()) {
      newErrors.domain = intl.formatMessage({ id: 'service.domainRequired' });
    }
    if (!formData.cluster) {
      newErrors.cluster = intl.formatMessage({ id: 'service.clusterRequired' });
    }

    setErrors(newErrors);
    return Object.keys(newErrors).length === 0;
  };

  const isValidUrl = (str: string): boolean => {
    try {
      new URL(str);
      return true;
    } catch {
      return false;
    }
  };

  const handleSubmit = () => {
    if (!validate()) {
      return;
    }
    onSubmit(formData);
    setFormData({
      name: '',
      repo: '',
      domain: '',
      cluster: '',
      description: '',
    });
  };

  if (!isOpen) {
    return null;
  }

  return (
    <div className={styles.modalOverlay} onClick={onClose}>
      <div className={styles.modalContent} onClick={(e) => e.stopPropagation()}>
        <div className={styles.modalHeader}>
          <h2 className={styles.modalTitle}>
            <FormattedMessage id="service.addService" />
          </h2>
          <button className={styles.modalClose} onClick={onClose}>
            ×
          </button>
        </div>

        <div className={styles.modalBody}>
          <div className={styles.formGroup}>
            <label className={styles.formLabel}>
              <FormattedMessage id="service.name" />
              <span className={styles.required}>*</span>
            </label>
            <input
              className={`${styles.formInput} ${
                errors.name ? styles.inputError : ''
              }`}
              type="text"
              value={formData.name}
              onChange={(e) => handleChange('name', e.target.value)}
              placeholder={intl.formatMessage({
                id: 'service.namePlaceholder',
              })}
            />
            {errors.name && (
              <span className={styles.errorText}>{errors.name}</span>
            )}
          </div>

          <div className={styles.formGroup}>
            <label className={styles.formLabel}>
              <FormattedMessage id="service.repo" />
              <span className={styles.required}>*</span>
            </label>
            <input
              className={`${styles.formInput} ${
                errors.repo ? styles.inputError : ''
              }`}
              type="text"
              value={formData.repo}
              onChange={(e) => handleChange('repo', e.target.value)}
              placeholder={intl.formatMessage({
                id: 'service.repoPlaceholder',
              })}
            />
            {errors.repo && (
              <span className={styles.errorText}>{errors.repo}</span>
            )}
          </div>

          <div className={styles.formGroup}>
            <label className={styles.formLabel}>
              <FormattedMessage id="service.domain" />
              <span className={styles.required}>*</span>
            </label>
            <input
              className={`${styles.formInput} ${
                errors.domain ? styles.inputError : ''
              }`}
              type="text"
              value={formData.domain}
              onChange={(e) => handleChange('domain', e.target.value)}
              placeholder={intl.formatMessage({
                id: 'service.domainPlaceholder',
              })}
            />
            {errors.domain && (
              <span className={styles.errorText}>{errors.domain}</span>
            )}
          </div>

          <div className={styles.formGroup}>
            <label className={styles.formLabel}>
              <FormattedMessage id="service.cluster" />
              <span className={styles.required}>*</span>
            </label>
            <select
              className={`${styles.formSelect} ${
                errors.cluster ? styles.inputError : ''
              }`}
              value={formData.cluster}
              onChange={(e) => handleChange('cluster', e.target.value)}
            >
              <option value="">
                {intl.formatMessage({ id: 'service.selectCluster' })}
              </option>
              {clusters.map((cluster) => (
                <option key={cluster.name} value={cluster.name}>
                  {cluster.name} ({cluster.tags.join(', ')})
                </option>
              ))}
            </select>
            {errors.cluster && (
              <span className={styles.errorText}>{errors.cluster}</span>
            )}
          </div>

          <div className={styles.formGroup}>
            <label className={styles.formLabel}>
              <FormattedMessage id="service.description" />
            </label>
            <textarea
              className={styles.formTextarea}
              value={formData.description}
              onChange={(e) => handleChange('description', e.target.value)}
              placeholder={intl.formatMessage({
                id: 'service.descriptionPlaceholder',
              })}
              rows={3}
            />
          </div>
        </div>

        <div className={styles.modalFooter}>
          <button className={styles.modalCancel} onClick={onClose}>
            <FormattedMessage id="service.cancel" />
          </button>
          <button className={styles.modalSubmit} onClick={handleSubmit}>
            <FormattedMessage id="service.submit" />
          </button>
        </div>
      </div>
    </div>
  );
}
