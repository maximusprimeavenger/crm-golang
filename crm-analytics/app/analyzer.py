import pandas as pd
import pickle
from sklearn.cluster import KMeans
from sklearn.preprocessing import StandardScaler
from sklearn.linear_model import LinearRegression

def fetch_data(gateway_url):
    import requests
    response = requests.get(gateway_url)
    response.raise_for_status()
    df = pd.DataFrame(response.json())
    if 'date' in df.columns:
        df['date'] = pd.to_datetime(df['date'])
    return df

def cluster_customers(df, n_clusters=3, train_new=False):
    features = df[['total_sales','visits','last_purchase_days']].fillna(0)
    scaler = StandardScaler()
    X_scaled = scaler.fit_transform(features)
    
    if train_new:
        kmeans = KMeans(n_clusters=n_clusters, random_state=42)
        df['cluster'] = kmeans.fit_predict(X_scaled)
        with open("app/models/customer_cluster.pkl","wb") as f:
            pickle.dump(kmeans,f)
    else:
        with open("app/models/customer_cluster.pkl","rb") as f:
            kmeans = pickle.load(f)
        df['cluster'] = kmeans.predict(X_scaled)
    return df

def forecast_sales(df, train_new=False):
    if 'date' not in df.columns:
        raise ValueError("DataFrame должен содержать колонку 'date'")
    
    df_grouped = df.groupby('date').total_sales.sum().reset_index()
    df_grouped['days'] = (df_grouped['date'] - df_grouped['date'].min()).dt.days
    
    X = df_grouped[['days']]
    y = df_grouped['total_sales']
    
    if train_new:
        model = LinearRegression()
        model.fit(X, y)
        with open("app/models/sales_model.pkl","wb") as f:
            pickle.dump(model,f)
    else:
        with open("app/models/sales_model.pkl","rb") as f:
            model = pickle.load(f)
    
    df_grouped['predicted_sales'] = model.predict(X)
    return df_grouped
