# Secret Manager - Bulk create and delete
1. Prepare a yaml file with secrets in below format
    ```yaml
    project: <GCP Project ID>
    secrets:
      - name: <secret1 name>
        value: <secret1 value>
        labels:
          <lable1_name>: <label1_value>
          <lable2_name>: <label2_value>
      - name: <secret2 name>
        value: <secret2 value>
        labels:
          <lable1_name>: <label1_value>
    ```
2. execute programe gce_sm.exe
    ```sh
    # Create Secrets
    gce_sm.exe bulk-create <path to secret yaml>
    # Delete Secrets
    gce_sm.exe bulk-delete <path to secret yaml>
    ```