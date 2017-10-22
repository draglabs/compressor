<?xml version="1.0" encoding="UTF-8" standalone="no" ?>
<!DOCTYPE xmeml>
<xmeml version="5">

  <project>
    <name>40 seconds for DL clip_01.xml</name>    // JamName.xml
    <children>
      <sequence id="sequence-1">
        <duration>112580</duration>        //Total Duration in "Frames"
        <rate>
          <timebase>30</timebase>          //There are 30 frames per second
          <ntsc>FALSE</ntsc>
        </rate>
        <name>40 seconds for DL clip_01.xml</name>  // JamName.x
        <media>
          <video>
            <format>
              <samplecharacteristics>
                <rate>
                  <timebase>30</timebase>
                  <ntsc>FALSE</ntsc>
                </rate>
                <width>1920</width>
                <height>1080</height>
                <anamorphic>FALSE</anamorphic>
                <pixelaspectratio>Square</pixelaspectratio>
                <fielddominance>upper</fielddominance>
                <colordepth>24</colordepth>
              </samplecharacteristics>
            </format>
            <track>
              <enabled>TRUE</enabled>
              <locked>FALSE</locked>
            </track>
          </video>
          <audio>
            <format>
              <samplecharacteristics>
                <depth>32</depth>
                <samplerate>48000</samplerate>
              </samplecharacteristics>
            </format>
            <outputs>
              <group>
                <index>1</index>
                <numchannels>1</numchannels>
                <downmix>0</downmix>
                <channel>
                  <index>1</index>
                </channel>
              </group>
              <group>
                <index>2</index>
                <numchannels>1</numchannels>
                <downmix>0</downmix>
                <channel>
                  <index>2</index>
                </channel>
              </group>
            </outputs>
            <track>
              <enabled>TRUE</enabled>
              <locked>FALSE</locked>
              <clipitem id="clipitem-0">
                <name>TASCAM_0003 48000 1</name>         // contributer1 name
                <enabled>TRUE</enabled>
                <duration>985</duration>                 // duration in "Frames"
                <start>7</start>                         // Start time in frames
                <end>992</end>                           // end time in "Frames"
                <in>238</in>                             // Same as Start time
                <out>1224</out>                          // Same as end time
                <file id="0">                            // Zero for the first track ... 1 for the second
                  <name>TASCAM_0003 48000 1.wav</name>   // name of CAF file
                  <pathurl>/sources/name-of-audio-file.caf</pathurl>
                  <rate>
                    <timebase>30</timebase>
                    <ntsc>FALSE</ntsc>
                  </rate>
                  <duration>1224</duration>
                  <media>
                    <audio>
                      <samplecharacteristics>
                        <depth>32</depth>
                        <samplerate>48000</samplerate>
                      </samplecharacteristics>
                    </audio>
                  </media>
                </file>
                <sourcetrack>
                  <mediatype>audio</mediatype>
                  <trackindex>1</trackindex>
                </sourcetrack>
                <channelcount>2</channelcount>
              </clipitem>
              <outputchannelindex>25</outputchannelindex>
            </track>
            <track>
              <enabled>TRUE</enabled>
              <locked>FALSE</locked>
              <clipitem id="clipitem-1">
                <name>TASCAM_0004 48000 1</name>
                <enabled>TRUE</enabled>
                <duration>1176</duration>
                <start>0</start>
                <end>1176</end>
                <in>215</in>
                <out>1392</out>
                <file id="1">
                  <name>TASCAM_0004 48000 1.wav</name>
                  <pathurl>file://localhost/Users/david/Documents/XMLTest/40%20seconds%20for%20DL%20clip_01_CopiedFiles/TASCAM_0004%2048000%201.wav</pathurl>
                  <rate>
                    <timebase>30</timebase>
                    <ntsc>FALSE</ntsc>
                  </rate>
                  <duration>111791</duration>
                  <media>
                    <audio>
                      <samplecharacteristics>
                        <depth>32</depth>
                        <samplerate>48000</samplerate>
                      </samplecharacteristics>
                    </audio>
                  </media>
                </file>
                <sourcetrack>
                  <mediatype>audio</mediatype>
                  <trackindex>1</trackindex>
                </sourcetrack>
                <channelcount>2</channelcount>
              </clipitem>
              <outputchannelindex>25</outputchannelindex>
            </track>
            <track>
              <enabled>TRUE</enabled>
              <locked>FALSE</locked>
              <clipitem id="clipitem-2">
                <name>TASCAM_0005 48000 1</name>
                <enabled>TRUE</enabled>
                <duration>1315</duration>
                <start>0</start>
                <end>1315</end>
                <in>190</in>
                <out>1505</out>
                <file id="2">
                  <name>TASCAM_0005 48000 1.wav</name>
                  <pathurl>file://localhost/Users/david/Documents/XMLTest/40%20seconds%20for%20DL%20clip_01_CopiedFiles/TASCAM_0005%2048000%201.wav</pathurl>
                  <rate>
                    <timebase>30</timebase>
                    <ntsc>FALSE</ntsc>
                  </rate>
                  <duration>68999</duration>
                  <media>
                    <audio>
                      <samplecharacteristics>
                        <depth>32</depth>
                        <samplerate>48000</samplerate>
                      </samplecharacteristics>
                    </audio>
                  </media>
                </file>
                <sourcetrack>
                  <mediatype>audio</mediatype>
                  <trackindex>1</trackindex>
                </sourcetrack>
                <channelcount>2</channelcount>
              </clipitem>
              <clipitem id="clipitem-3">
                <name>App that could do that_audio 48000 1</name>
                <enabled>TRUE</enabled>
                <duration>177</duration>
                <start>1123</start>
                <end>1301</end>
                <in>0</in>
                <out>177</out>
                <file id="3">
                  <name>App that could do that_audio 48000 1.wav</name>
                  <pathurl>file://localhost/Users/david/Documents/XMLTest/40%20seconds%20for%20DL%20clip_01_CopiedFiles/App%20that%20could%20do%20that_audio%2048000%201.wav</pathurl>
                  <rate>
                    <timebase>30</timebase>
                    <ntsc>FALSE</ntsc>
                  </rate>
                  <duration>177</duration>
                  <media>
                    <audio>
                      <samplecharacteristics>
                        <depth>32</depth>
                        <samplerate>48000</samplerate>
                      </samplecharacteristics>
                    </audio>
                  </media>
                </file>
                <sourcetrack>
                  <mediatype>audio</mediatype>
                  <trackindex>1</trackindex>
                </sourcetrack>
                <channelcount>2</channelcount>
              </clipitem>
              <outputchannelindex>25</outputchannelindex>
            </track>
            <track>
              <enabled>FALSE</enabled>
              <locked>FALSE</locked>
              <clipitem id="clipitem-4">
                <name>What Keeps me up_1_audio 48000 1</name>
                <enabled>TRUE</enabled>
                <duration>50</duration>
                <start>0</start>
                <end>50</end>
                <in>3164</in>
                <out>3214</out>
                <file id="4">
                  <name>What Keeps me up_1_audio 48000 1.wav</name>
                  <pathurl>file://localhost/Users/david/Documents/XMLTest/40%20seconds%20for%20DL%20clip_01_CopiedFiles/What%20Keeps%20me%20up_1_audio%2048000%201.wav</pathurl>
                  <rate>
                    <timebase>30</timebase>
                    <ntsc>FALSE</ntsc>
                  </rate>
                  <duration>8270</duration>
                  <media>
                    <audio>
                      <samplecharacteristics>
                        <depth>32</depth>
                        <samplerate>48000</samplerate>
                      </samplecharacteristics>
                    </audio>
                  </media>
                </file>
                <sourcetrack>
                  <mediatype>audio</mediatype>
                  <trackindex>1</trackindex>
                </sourcetrack>
                <channelcount>2</channelcount>
              </clipitem>
              <clipitem id="clipitem-5">
                <name>What Keeps me up_1_audio 48000 1</name>
                <enabled>TRUE</enabled>
                <duration>128</duration>
                <start>45</start>
                <end>173</end>
                <in>3292</in>
                <out>3420</out>
                <file id="4">
                  <media>
                    <audio>
                      <samplecharacteristics>
                        <depth>32</depth>
                        <samplerate>48000</samplerate>
                      </samplecharacteristics>
                    </audio>
                  </media>
                </file>
                <sourcetrack>
                  <mediatype>audio</mediatype>
                  <trackindex>1</trackindex>
                </sourcetrack>
                <channelcount>2</channelcount>
              </clipitem>
              <clipitem id="clipitem-6">
                <name>What Keeps me up_1_audio 48000 1</name>
                <enabled>TRUE</enabled>
                <duration>279</duration>
                <start>173</start>
                <end>452</end>
                <in>3634</in>
                <out>3913</out>
                <file id="4">
                  <media>
                    <audio>
                      <samplecharacteristics>
                        <depth>32</depth>
                        <samplerate>48000</samplerate>
                      </samplecharacteristics>
                    </audio>
                  </media>
                </file>
                <sourcetrack>
                  <mediatype>audio</mediatype>
                  <trackindex>1</trackindex>
                </sourcetrack>
                <channelcount>2</channelcount>
              </clipitem>
              <clipitem id="clipitem-7">
                <name>What Keeps me up_1_audio 48000 1</name>
                <enabled>TRUE</enabled>
                <duration>214</duration>
                <start>452</start>
                <end>666</end>
                <in>4254</in>
                <out>4468</out>
                <file id="4">
                  <media>
                    <audio>
                      <samplecharacteristics>
                        <depth>32</depth>
                        <samplerate>48000</samplerate>
                      </samplecharacteristics>
                    </audio>
                  </media>
                </file>
                <sourcetrack>
                  <mediatype>audio</mediatype>
                  <trackindex>1</trackindex>
                </sourcetrack>
                <channelcount>2</channelcount>
              </clipitem>
              <clipitem id="clipitem-8">
                <name>What Keeps me up_1_audio 48000 1</name>
                <enabled>TRUE</enabled>
                <duration>401</duration>
                <start>666</start>
                <end>1068</end>
                <in>4741</in>
                <out>5143</out>
                <file id="4">
                  <media>
                    <audio>
                      <samplecharacteristics>
                        <depth>32</depth>
                        <samplerate>48000</samplerate>
                      </samplecharacteristics>
                    </audio>
                  </media>
                </file>
                <sourcetrack>
                  <mediatype>audio</mediatype>
                  <trackindex>1</trackindex>
                </sourcetrack>
                <channelcount>2</channelcount>
              </clipitem>
              <clipitem id="clipitem-9">
                <name>App that could do that_audio 48000 1</name>
                <enabled>TRUE</enabled>
                <duration>7</duration>
                <start>890</start>
                <end>898</end>
                <in>63</in>
                <out>71</out>
                <file id="3">
                  <media>
                    <audio>
                      <samplecharacteristics>
                        <depth>32</depth>
                        <samplerate>48000</samplerate>
                      </samplecharacteristics>
                    </audio>
                  </media>
                </file>
                <sourcetrack>
                  <mediatype>audio</mediatype>
                  <trackindex>1</trackindex>
                </sourcetrack>
                <channelcount>2</channelcount>
              </clipitem>
              <clipitem id="clipitem-10">
                <name>App that could do that_audio 48000 1</name>
                <enabled>TRUE</enabled>
                <duration>7</duration>
                <start>890</start>
                <end>897</end>
                <in>64</in>
                <out>71</out>
                <file id="3">
                  <media>
                    <audio>
                      <samplecharacteristics>
                        <depth>32</depth>
                        <samplerate>48000</samplerate>
                      </samplecharacteristics>
                    </audio>
                  </media>
                </file>
                <sourcetrack>
                  <mediatype>audio</mediatype>
                  <trackindex>1</trackindex>
                </sourcetrack>
                <channelcount>2</channelcount>
              </clipitem>
              <clipitem id="clipitem-11">
                <name>App that could do that_audio 48000 1</name>
                <enabled>TRUE</enabled>
                <duration>7</duration>
                <start>892</start>
                <end>900</end>
                <in>63</in>
                <out>71</out>
                <file id="3">
                  <media>
                    <audio>
                      <samplecharacteristics>
                        <depth>32</depth>
                        <samplerate>48000</samplerate>
                      </samplecharacteristics>
                    </audio>
                  </media>
                </file>
                <sourcetrack>
                  <mediatype>audio</mediatype>
                  <trackindex>1</trackindex>
                </sourcetrack>
                <channelcount>2</channelcount>
              </clipitem>
              <clipitem id="clipitem-12">
                <name>What Keeps me up_1_audio 48000 1</name>
                <enabled>TRUE</enabled>
                <duration>8</duration>
                <start>900</start>
                <end>908</end>
                <in>4064</in>
                <out>4073</out>
                <file id="4">
                  <media>
                    <audio>
                      <samplecharacteristics>
                        <depth>32</depth>
                        <samplerate>48000</samplerate>
                      </samplecharacteristics>
                    </audio>
                  </media>
                </file>
                <sourcetrack>
                  <mediatype>audio</mediatype>
                  <trackindex>1</trackindex>
                </sourcetrack>
                <channelcount>2</channelcount>
              </clipitem>
              <clipitem id="clipitem-13">
                <name>the furetures now2 48000 1</name>
                <enabled>TRUE</enabled>
                <duration>100</duration>
                <start>1082</start>
                <end>1183</end>
                <in>0</in>
                <out>100</out>
                <file id="5">
                  <name>the furetures now2 48000 1.wav</name>
                  <pathurl>file://localhost/Users/david/Documents/XMLTest/40%20seconds%20for%20DL%20clip_01_CopiedFiles/the%20furetures%20now2%2048000%201.wav</pathurl>
                  <rate>
                    <timebase>30</timebase>
                    <ntsc>FALSE</ntsc>
                  </rate>
                  <duration>93</duration>
                  <media>
                    <audio>
                      <samplecharacteristics>
                        <depth>32</depth>
                        <samplerate>48000</samplerate>
                      </samplecharacteristics>
                    </audio>
                  </media>
                </file>
                <sourcetrack>
                  <mediatype>audio</mediatype>
                  <trackindex>1</trackindex>
                </sourcetrack>
                <channelcount>2</channelcount>
              </clipitem>
              <clipitem id="clipitem-14">
                <name>The future is now 48000 1</name>
                <enabled>TRUE</enabled>
                <duration>111</duration>
                <start>1446</start>
                <end>1558</end>
                <in>0</in>
                <out>111</out>
                <file id="6">
                  <name>The future is now 48000 1.wav</name>
                  <pathurl>file://localhost/Users/david/Documents/XMLTest/40%20seconds%20for%20DL%20clip_01_CopiedFiles/The%20future%20is%20now%2048000%201.wav</pathurl>
                  <rate>
                    <timebase>30</timebase>
                    <ntsc>FALSE</ntsc>
                  </rate>
                  <duration>103</duration>
                  <media>
                    <audio>
                      <samplecharacteristics>
                        <depth>32</depth>
                        <samplerate>48000</samplerate>
                      </samplecharacteristics>
                    </audio>
                  </media>
                </file>
                <sourcetrack>
                  <mediatype>audio</mediatype>
                  <trackindex>1</trackindex>
                </sourcetrack>
                <channelcount>2</channelcount>
              </clipitem>
              <outputchannelindex>25</outputchannelindex>
            </track>
            <track>
              <enabled>TRUE</enabled>
              <locked>FALSE</locked>
              <clipitem id="clipitem-15">
                <name>ANW2780_07_Secret-Somewhere</name>
                <enabled>TRUE</enabled>
                <duration>1070</duration>
                <start>5</start>
                <end>1076</end>
                <in>190</in>
                <out>1260</out>
                <file id="7">
                  <name>ANW2780_07_Secret-Somewhere.wav</name>
                  <pathurl>file://localhost/Users/david/Documents/XMLTest/40%20seconds%20for%20DL%20clip_01_CopiedFiles/ANW2780_07_Secret-Somewhere.wav</pathurl>
                  <rate>
                    <timebase>30</timebase>
                    <ntsc>FALSE</ntsc>
                  </rate>
                  <duration>5587</duration>
                  <media>
                    <audio>
                      <samplecharacteristics>
                        <depth>32</depth>
                        <samplerate>48000</samplerate>
                      </samplecharacteristics>
                    </audio>
                  </media>
                </file>
                <sourcetrack>
                  <mediatype>audio</mediatype>
                  <trackindex>1</trackindex>
                </sourcetrack>
                <channelcount>2</channelcount>
              </clipitem>
              <clipitem id="clipitem-16">
                <name>ANW2780_07_Secret-Somewhere</name>
                <enabled>TRUE</enabled>
                <duration>245</duration>
                <start>1076</start>
                <end>1321</end>
                <in>0</in>
                <out>245</out>
                <file id="8">
                  <name>ANW2780_07_Secret-Somewhere_merged.wav</name>
                  <pathurl>file://localhost/Users/david/Documents/XMLTest/40%20seconds%20for%20DL%20clip_01_CopiedFiles/ANW2780_07_Secret-Somewhere_merged.wav</pathurl>
                  <rate>
                    <timebase>30</timebase>
                    <ntsc>FALSE</ntsc>
                  </rate>
                  <duration>245</duration>
                  <media>
                    <audio>
                      <samplecharacteristics>
                        <depth>32</depth>
                        <samplerate>48000</samplerate>
                      </samplecharacteristics>
                    </audio>
                  </media>
                </file>
                <sourcetrack>
                  <mediatype>audio</mediatype>
                  <trackindex>1</trackindex>
                </sourcetrack>
                <channelcount>2</channelcount>
              </clipitem>
              <outputchannelindex>25</outputchannelindex>
            </track>
            <track>
              <enabled>TRUE</enabled>
              <locked>FALSE</locked>
              <clipitem id="clipitem-17">
                <name>App that could do that_audio 48000 1</name>
                <enabled>TRUE</enabled>
                <duration>177</duration>
                <start>1062</start>
                <end>1240</end>
                <in>0</in>
                <out>177</out>
                <file id="3">
                  <media>
                    <audio>
                      <samplecharacteristics>
                        <depth>32</depth>
                        <samplerate>48000</samplerate>
                      </samplecharacteristics>
                    </audio>
                  </media>
                </file>
                <sourcetrack>
                  <mediatype>audio</mediatype>
                  <trackindex>1</trackindex>
                </sourcetrack>
                <channelcount>2</channelcount>
              </clipitem>
              <outputchannelindex>25</outputchannelindex>
            </track>
            <track>
              <enabled>TRUE</enabled>
              <locked>FALSE</locked>
              <clipitem id="clipitem-18">
                <name>What Keeps me up_1_audio 48000 1</name>
                <enabled>TRUE</enabled>
                <duration>1179</duration>
                <start>17</start>
                <end>1197</end>
                <in>0</in>
                <out>1179</out>
                <file id="9">
                  <name>What Keeps me up_1_audio 48000 1_merged.wav</name>
                  <pathurl>file://localhost/Users/david/Documents/XMLTest/40%20seconds%20for%20DL%20clip_01_CopiedFiles/What%20Keeps%20me%20up_1_audio%2048000%201_merged.wav</pathurl>
                  <rate>
                    <timebase>30</timebase>
                    <ntsc>FALSE</ntsc>
                  </rate>
                  <duration>1179</duration>
                  <media>
                    <audio>
                      <samplecharacteristics>
                        <depth>32</depth>
                        <samplerate>48000</samplerate>
                      </samplecharacteristics>
                    </audio>
                  </media>
                </file>
                <sourcetrack>
                  <mediatype>audio</mediatype>
                  <trackindex>1</trackindex>
                </sourcetrack>
                <channelcount>2</channelcount>
              </clipitem>
              <outputchannelindex>25</outputchannelindex>
            </track>
          </audio>
        </media>
        <timecode>
          <rate>
            <timebase>30</timebase>
            <ntsc>FALSE</ntsc>
          </rate>
          <frame>0</frame>
          <displayformat>NDF</displayformat>
        </timecode>
      </sequence>
    </children>
  </project>

</xmeml>