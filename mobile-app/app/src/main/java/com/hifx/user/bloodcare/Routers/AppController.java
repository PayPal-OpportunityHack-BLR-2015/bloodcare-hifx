package com.hifx.user.bloodcare.Routers;

import android.app.Application;
import android.content.Context;
import android.database.SQLException;
import android.graphics.Typeface;
import android.net.ConnectivityManager;
import android.net.NetworkInfo;
import android.os.AsyncTask;
import android.os.SystemClock;
import android.text.TextUtils;
import android.util.Log;
import android.widget.Toast;

import com.android.volley.AuthFailureError;
import com.android.volley.Request;
import com.android.volley.RequestQueue;
import com.android.volley.Response;
import com.android.volley.VolleyError;
import com.android.volley.VolleyLog;
import com.android.volley.toolbox.ImageLoader;
import com.android.volley.toolbox.JsonArrayRequest;
import com.android.volley.toolbox.JsonObjectRequest;
import com.android.volley.toolbox.Volley;
import com.hifx.user.bloodcare.Constants;
import com.hifx.user.bloodcare.DatabaseHandler.DatabaseHandler;

import org.json.JSONArray;
import org.json.JSONException;
import org.json.JSONObject;

import java.util.HashMap;
import java.util.Map;

public class AppController extends Application {
	JSONArray channelArray = null;
	public Object main;

	public static DatabaseHandler databasehandler;
	public static final int TIMEOUTINTERVAL = 60000;
	Context alertcontext;
	public float zoomvalue = 0;
	public boolean pushnotification = true;
	public static float detail_image_size;
	Object customlist = null;
	JsonResponseInterface responseInterface = null;
	public static Typeface font;
	public static String GCMregister = null;
	public int fontposition = 0;
	public static int numColomns = 2;
	public JSONObject credentials = null;
	public String id_token = null;
	public String user_name = null;
	public String profileImage = "";
	public String email = null;
	public long expiry=0;
	public String user_id = null;
    public static String commentnumbr = "";
    public static String[] tileColor = { "#197c8f", "#525f69", "#621e43", "#ac3538",
            "#2d2d2d", "#2e4346", "#3d692e", "#022951","#197c8f","#525f69" };

	public Typeface setTypeface() {
		if (font == null) {
			font = Typeface.createFromAsset(getAssets(), "fonts/Manorama.otf");
		}

		return font;
	}

	public void setCustomListView(Object customlist) {
		this.customlist = customlist;

	}

	public void clearCache() {
		new clearCacheFromDb().execute();

	}

	public Object getCustomListView() {
		return this.customlist;

	}

	public void setZoomValue(float f) {
		this.zoomvalue = f;

	}

	public static final String TAG = AppController.class.getSimpleName();

	public float getZoomValue() {
		return this.zoomvalue;

	}

	public Context getContext() {
		return this.alertcontext;
	}

	public void setContext(Context context) {
		// Log.i("context is", this.alertcontext.toString());
		this.alertcontext = context;
	}

	public boolean CheckNetWorkConnection() {
		ConnectivityManager cm = (ConnectivityManager) getApplicationContext()
				.getSystemService(Context.CONNECTIVITY_SERVICE);

		NetworkInfo activeNetwork = cm.getActiveNetworkInfo();

		if (activeNetwork != null && activeNetwork.isAvailable()
				&& activeNetwork.isConnected()) {

			return true;
		} else {

			return false;

		}

	}

	private class clearCacheFromDb extends AsyncTask<String, Integer, String> {

		@Override
		protected String doInBackground(String... params) {
			try {
				databasehandler = DatabaseHandler
						.getInstance(getApplicationContext());
				if (databasehandler != null) {
					databasehandler.clearCache();
				}
				

				return "Cache cleared successfully";
			} catch (Exception e) {
				return "Cache cleared successfully";

			}

		}

		@Override
		protected void onProgressUpdate(Integer... progress) {
		}

		@Override
		protected void onPostExecute(String result) {
			Toast.makeText(getApplicationContext(), result, Toast.LENGTH_SHORT)
					.show();

		}

	}

	private RequestQueue mRequestQueue;
	private ImageLoader mImageLoader;

	private static AppController mInstance;

	@Override
	public void onCreate() {
		super.onCreate();
		mInstance = this;
	}

	public static synchronized AppController getInstance() {
		return mInstance;
	}

	public RequestQueue getRequestQueue() {
		if (mRequestQueue == null) {
			mRequestQueue = Volley.newRequestQueue(getApplicationContext());
		}

		return mRequestQueue;
	}
//
//	public ImageLoader getImageLoader() {
//		getRequestQueue();
//		if (mImageLoader == null) {
//			mImageLoader = new ImageLoader(this.mRequestQueue,
//					new LruBitmapCache());
//		}
//		return this.mImageLoader;
//	}

	public <T> void addToRequestQueue(Request<T> req, String tag) {
		// set the default tag if tag is empty
		req.setTag(TextUtils.isEmpty(tag) ? TAG : tag);
		getRequestQueue().add(req);
	}

	public <T> void addToRequestQueue(Request<T> req) {
		req.setTag(TAG);
		getRequestQueue().add(req);
	}

	public void cancelPendingRequests(Object tag) {
		if (mRequestQueue != null) {
			mRequestQueue.cancelAll(tag);
		}
	}




	public void FetchjsonFromurl(String url, final String code,
			final String pagenumber, final boolean saveToDB,
			final String parentcode, final String classfunctionname,
			final String retainorDelete, final JsonResponseInterface callback) {
		Log.e("url", url + "fetchjsonfromurl");
		String tag_json_obj = "json_obj_req";
		JsonObjectRequest jsonObjReq = new JsonObjectRequest(url, null,
				new Response.Listener<JSONObject>() {

					@SuppressWarnings("unchecked")
					@Override
					public void onResponse(JSONObject object) {
						Log.i("obje", object.toString());
						HashMap<String, String> hash = new HashMap<String, String>();
						hash.put(Constants.code, code);
						hash.put(Constants.pagenumber, pagenumber);
						hash.put(Constants.newsdata, object.toString());

						String expiry = null;
						long sysTime = SystemClock.elapsedRealtime() / 1000;
						try {
						if (object.has("header")) {
							JSONObject header = object
							.getJSONObject("header");
							if(header.has("expiryTime")){
								expiry = String.valueOf(sysTime
                                        + Long.valueOf(header
                                        .get("expiryTime").toString()));
							}

							else {
								expiry = String.valueOf(sysTime
                                        + String.valueOf(24 * 60 * 60));
							}

						} else {
							expiry = String.valueOf(sysTime
                                    + String.valueOf(24 * 60 * 60));
						}
						} catch (JSONException e) {
							// TODO Auto-generated catch block
							e.printStackTrace();
						}
						hash.put(Constants.expiry, expiry);
						hash.put(Constants.logTime, String.valueOf(sysTime));
						responseInterface = callback;
						responseInterface
								.dataFromApi(object, classfunctionname);
						if (saveToDB) {

							hash.put(Constants.parentCode, parentcode);

							if (retainorDelete.equalsIgnoreCase("delete")) {
								new deleteParentCode().execute(parentcode);
							}
							// meeranew
							new InsertDataToDatabase().execute(hash);
						}

					}
				}, new Response.ErrorListener() {

					@Override
					public void onErrorResponse(VolleyError error) {
						Log.e("eroorrr", "error	" + error.getMessage() + "****"
                                + error.networkResponse);

						VolleyLog.d(TAG, "Error: " + error.getMessage());
						int statuscode = -1;
						if (error != null) {
							if (error.networkResponse != null) {
								statuscode = error.networkResponse.statusCode;

							}

						}
						responseInterface = callback;
						responseInterface.ErrorResponseVolley(statuscode,
								classfunctionname);

					}
				}) {
			@Override
			public Map<String, String> getHeaders() throws AuthFailureError {
				Map<String, String> header = new HashMap<String, String>();
				if(classfunctionname.contains("getSpamDialog")){
				if (credentials != null
						&& credentials.length() > 0) {
					
					header.put("Authorization", "Bearer " + id_token);

				}
                    else{
                     header.put("User-Agent", "Android");
                }
				}
                else {
                    header.put("User-Agent", "Android");
                }
                return header;
			}
		};
		// Adding request to request queue
		AppController.getInstance().addToRequestQueue(jsonObjReq, tag_json_obj);
	}

    public void FetchjsonVideo(String url, final String code,
                               final String pagenumber, final boolean saveToDB,
                               final String parentcode, final String classfunctionname,
                               final String retainorDelete, final JsonResponseInterfaceArray callback) {

        Log.e("videeoooo","videoooo");
        String tag_json_arry = "json_array_req";


        JsonArrayRequest req = new JsonArrayRequest(url,
                new Response.Listener<JSONArray>() {
                    @Override
                    public void onResponse(JSONArray response) {

                        Log.i("res",response.toString());
                        if (response != null && response.length() > 0) {

                            try {
                                    HashMap<String, String> hash = new HashMap<String, String>();
                                    hash.put(Constants.code, code);
                                    hash.put(Constants.pagenumber, pagenumber);
                                    hash.put(Constants.newsdata, response.toString());

                                    String expiry = null;
                                    long sysTime = SystemClock.elapsedRealtime() / 1000;

                                        expiry = String.valueOf(sysTime
                                                + String.valueOf(24 * 60 * 60));



                                    hash.put(Constants.expiry, expiry);
                                    hash.put(Constants.logTime, String.valueOf(sysTime));
                                    callback
                                            .dataFromApi(response, classfunctionname);
                                    if (saveToDB) {

                                        hash.put(Constants.parentCode, parentcode);

                                        if (retainorDelete.equalsIgnoreCase("delete")) {
                                            new deleteParentCode().execute(parentcode);
                                        }
                                        // meeranew
                                        new InsertDataToDatabase().execute(hash);
                                    }

                            } catch (Exception e) {
                                // TODO Auto-generated catch block
                                e.printStackTrace();
                            }

                        }

                    }
                }, new Response.ErrorListener() {
            @Override
            public void onErrorResponse(VolleyError error) {
                VolleyLog.d(TAG, "Error: " + error.getMessage());
                int statuscode = -1;
                if (error != null) {
                    if (error.networkResponse != null) {
                        statuscode = error.networkResponse.statusCode;

                    }

                }
                callback.ErrorResponseVolley(statuscode,
                        classfunctionname);
            }
        });

        AppController.getInstance().addToRequestQueue(req, tag_json_arry);

    }





//	public void Fetchjsonget(String url, final Map<String, String> header,
//			final String classfunctionname, final String calledfunction,
//
//			final JsonResponseInterface callback) {
//		Log.i("url", url);
//
//		JsonObjectRequest myRequest = new JsonObjectRequest(Request.Method.GET,
//				url, null,
//
//				new Response.Listener<JSONObject>() {
//					@Override
//					public void onResponse(JSONObject response) {
//						responseInterface = callback;
//						responseInterface.dataFromApi(response,
//								classfunctionname);
//					}
//				}, new Response.ErrorListener() {
//					@Override
//					public void onErrorResponse(VolleyError error) {
//						Log.i("error in gettingg", "error in gettingg");
//						int statuscode=-1;
//						if (error != null) {
//							if (error.networkResponse != null) {
//								 statuscode = error.networkResponse.statusCode;
//
//
//							}
//
//						}
//						responseInterface = callback;
//						responseInterface.ErrorResponseVolley(
//								statuscode, calledfunction);
//					}
//				}) {
//            @Override
//            protected Response<JSONObject> parseNetworkResponse(NetworkResponse response) {
//                try {
//                    // solution 1:
//                    String jsonString = new String(response.data, "UTF-8");
//
//                    return Response.success(new JSONObject(jsonString),
//                            HttpHeaderParser.parseCacheHeaders(response));
//                } catch (UnsupportedEncodingException e) {
//                    return Response.error(new ParseError(e));
//                } catch (JSONException je) {
//                    return Response.error(new ParseError(je));
//                }
//            }
//			@Override
//			public Map<String, String> getHeaders() throws AuthFailureError {
//
//
//				return header;
//			}
//		};
//
//		// Adding request to request queue
//		AppController.getInstance().addToRequestQueue(myRequest);
//
//	}



	private class deleteParentCode extends AsyncTask<String, Integer, String> {

		@Override
		protected String doInBackground(String... params) {

			// if (!params[0].equals("latest")) {
			//
			// if (ImageLoader.getInstance().isInited()) {
			// ImageLoader.getInstance().clearDiscCache();
			// } else {
			// ImageLoader.getInstance().init(
			// ImageLoaderConfiguration
			// .createDefault(getApplicationContext()));
			// ImageLoader.getInstance().clearDiscCache();
			// }
			//
			// }
			databasehandler = DatabaseHandler
					.getInstance(getApplicationContext());
			databasehandler.deleteParentcode(params[0]);
			return params[0];

		}

		@Override
		protected void onProgressUpdate(Integer... progress) {
		}

		@Override
		protected void onPostExecute(String result) {

		}

	}

	private class InsertDataToDatabase extends
            AsyncTask<HashMap<String, String>, Integer, String> {

		@Override
		protected String doInBackground(HashMap<String, String>... params) {
			HashMap<String, String> arrayListVal = params[0];
			String code = arrayListVal.get(Constants.code);
			String pagenumber = arrayListVal.get(Constants.pagenumber);
			String data = arrayListVal.get(Constants.newsdata);
			String date = arrayListVal.get(Constants.expiry);

			String parentCode = arrayListVal.get(Constants.parentCode);
			long logtimeDate = SystemClock.elapsedRealtime() / 1000;

			// Toast.makeText(getApplicationContext(),
			// logtimeDate+"**"+logtimeDate.getTime(),
			// Toast.LENGTH_SHORT).show();
			databasehandler = DatabaseHandler
					.getInstance(getApplicationContext());
			try {
				if (databasehandler != null) {
					databasehandler.insertNews(code,
							Integer.valueOf(pagenumber), data, date, 0,
							String.valueOf(logtimeDate), parentCode);
				}
			} catch (SQLException e) {
				e.printStackTrace();
			}
			return parentCode;

		}

		protected void onProgressUpdate(Integer... progress) {
		}

		protected void onPostExecute(String result) {
		}

	}

}
