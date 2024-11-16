# data-analysis-python\src\get_all_stock_info\get_all_stock_info_service.py
import investpy

def get_all_stock_info():
    try:
        # 全世界の株式情報を取得
        stocks = investpy.stocks.get_stocks()
        # 利用可能な列を確認
        available_columns = stocks.columns
        
        stock_info_list = []
        for _, row in stocks.iterrows():
            stock_info = {
                'country': row['country'] if 'country' in available_columns else '',
                'symbol': row['symbol'] if 'symbol' in available_columns else '',
                'name': row['name'] if 'name' in available_columns else '',
                'full_name': row['full_name'] if 'full_name' in available_columns else '',
                'isin': row['isin'] if 'isin' in available_columns else '',
                'currency': row['currency'] if 'currency' in available_columns else '',
                'stock_exchange': row['stock_exchange'] if 'stock_exchange' in available_columns else '',
                'sector': row['sector'] if 'sector' in available_columns else '',
                'industry': row['industry'] if 'industry' in available_columns else ''
            }
            stock_info_list.append(stock_info)
        return stock_info_list
    except Exception as e:
        print(f"Error fetching stock info: {e}")
        return []
